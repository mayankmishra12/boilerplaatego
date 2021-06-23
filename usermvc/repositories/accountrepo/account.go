package accountrepo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"net/url"
	Config "usermvc/config"
	"usermvc/entity"
	"usermvc/model"
	logger2 "usermvc/utility/logger"
)

type AccountRepo interface {
	//Create(context context.Context, accountDetailsRequest entity.AccountDetails) ()
	Insert(context context.Context, accountDetailsRequest entity.AccountDetails) (*model.AccountDetailsResponse, error)
}

type accountRepo struct {
	db *gorm.DB
}

func NewAccountRepo() *accountRepo {
	newDb, err := newDb()
	if err != nil {
		panic(err)
	}
	newDb.AutoMigrate(&entity.User{})
	return &accountRepo{
		db: newDb,
	}
}

func newDb() (*gorm.DB, error) {
	conf := Config.NewDbConfig()
	dsn := url.URL{
		User:     url.UserPassword(conf.User, conf.Password),
		Scheme:   "postgres",
		Host:     fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Path:     conf.DBName,
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}
	db, err := gorm.Open("postgres", dsn.String())
	if err != nil {
		return nil, err
	}
	return db, nil
}


func (r accountRepo) Create(context context.Context, account entity.AccountDetails) error {
	if err := r.db.Create(account).Error; err != nil {
		return err
	}
	return nil
}

func (r accountRepo) Insert(context context.Context, accountDetailsRequest entity.AccountDetails) (*model.AccountDetailsResponse, error)  {
	logger := logger2.GetLoggerWithContext(context)
	lead := accountDetailsRequest
	//var rows *sql.Rows
    fmt.Println("prting wiht the ")
	if (lead.Role == "Marketing Executive") && (lead.ConvertLeadToAccount) {

		sqlStatementl1 := `UPDATE CMS_LEADS_MASTER
						  	  SET 
						  	  masterstatus='Pending Approval'
						      WHERE 
					    	  leadid=$1`
		rows, err := r.db.DB().Query(sqlStatementl1, lead.LeadId)
		logger.Info()
		logger.Info("Updated Status to Pending Approval")

		if err != nil {
			logger.Info(err.Error())
			return nil, errors.New(err.Error())
			return nil, err
			//return events.APIGatewayProxyResponse{500, headers, nil, err.Error(), false}, nil
		}
		defer rows.Close()

		res, _ := json.Marshal(rows)
		fmt.Println(res)
		return  nil, nil
		//return events.APIGatewayProxyResponse{200, headers, nil, string(res), false}, nil

		// LEAD REJECTION MODULE
	} else if (lead.Role == "Managing Director") && (lead.Reject) {

		sqlStatementmdr1 := `UPDATE CMS_LEADS_MASTER
							SET 
							masterstatus='Rejected',
							comments=$1
	 						WHERE 
	  						leadid=$2`
		rows, err := r.db.DB().Query(sqlStatementmdr1, lead.Comments, lead.LeadId)
		logger.Info("Updated Status to Rejected")

		if err != nil {
			logger.Info(err.Error())
			return nil, err
			//return events.APIGatewayProxyResponse{500, headers, nil, err.Error(), false}, nil
		}
		defer rows.Close()

		res, _ := json.Marshal(rows)
		fmt.Println(res)
		return &model.AccountDetailsResponse{
			StatusCode: 200,
			Payload:    res,
		}, nil
		//return events.APIGatewayProxyResponse{200, headers, nil, string(res), false}, nil
		// APPROVAL BY MD
	} else if (lead.Role == "Managing Director") && (lead.ConvertLeadToAccount || lead.Approve) {

		sqlStatementacn1 := `UPDATE CMS_LEADS_MASTER
							  SET 
							  masterstatus='account Created',
					  		  accountid=(select floor(100 + random() * 899)::numeric)
					  		  WHERE 
					  		  leadid=$1`
		rows, err := r.db.DB().Query(sqlStatementacn1, lead.LeadId)
		logger.Info("account ID assigned")
		sqlStatementina1 := `INSERT INTO Accounts_master (
			accountid,
			accountname,
			accounttypeid,
			phone,
			email,
			createddate,
			createduserid,
			approxannualrev,
			website,
			productsegmentid,
			masterstatus,
			recordtypeid,
			shipping_continent,
			shipping_country,
			comments,
			aliases,
			isactive,
			otherinformation)
			SELECT
			accountid,
			accountname,
			accounttypeid,
			phone,
			email,
			createddate,
			createduserid ,
			approxannualrev,
			website,
			productsegmentid,
			masterstatus,
			recordtypeid,
			shipping_continent,
			shipping_country,
			comments,
			aliases,
			isactive,
			otherinformation
			FROM
			cms_leads_master
			WHERE leadid=$1`
		rows, err = r.db.DB().Query(sqlStatementina1, lead.LeadId)
		// Get Accountid from Lead Record
		// Set account status to Prospect in accounts_master
		sqlStatementstat1 := `UPDATE accounts_master
					 		 SET 
					 		 account_owner=u.username,
					 		 masterstatus='Prospect',
							 comments=$1
					 		 FROM accounts_master acc
					 		 INNER JOIN
					 		 CMS_LEADS_MASTER ld
					 		 ON ld.accountid = acc.accountid
					 		 INNER JOIN
					 		 userdetails_master u on u.userid=ld.createduserid
					 		 where ld.leadid=$2`

		rows, err = r.db.DB().Query(sqlStatementstat1, lead.Comments, lead.LeadId)
		sqlStatementapp1 := `UPDATE CMS_LEADS_MASTER
					  	  SET 
					  	  masterstatus='Appoved'
						  comments=$1	
					      WHERE 
					      leadid=$2`

		rows, err = r.db.DB().Query(sqlStatementapp1, lead.Comments, lead.LeadId)
		fmt.Println("Lead Status is set to Approved")
		sqlStatementcon1 := `insert into contacts_master(
			contactfirstname,
			contactlastname,
			contactemail,
			contactphone,
			contactmobilenumber,
			accountid,
			position,
			salutationid) 
			select
			contactfirstname,
			contactlastname,
			email,
			phone,
			contact_mobile,
			accountid,
			contact_position,
			contact_salutationid
			from
			cms_leads_master where leadid=$1`
		rows, err = r.db.DB().Query(sqlStatementcon1, lead.LeadId)
		fmt.Println("Lead Contact data is inserted into Contacts_Master Successfully")

		//Insert into accounts_billing_address_master
		sqlStatementabm1 := `insert into accounts_billing_address_master(
			accountid,
			billingid,
			street,
			city,
			stateprovince,
			postalcode,
			country)
			select
			ld.accountid,
			lba.billingid,
			lba.street,
			lba.city,
			lba.stateprovince,
			lba.postalcode,
			lba.country
			from
			cms_leads_billing_address_master lba
			inner join
			cms_leads_master ld on ld.leadid=lba.leadid
			where ld.leadid=$1`
		rows, err = r.db.DB().Query(sqlStatementabm1, lead.LeadId)
		fmt.Println("account Contact data is inserted into accounts_billing_address_master")
		//Insert into accounts_shipping_address_master
		sqlStatementasm1 := `insert into accounts_shipping_address_master(
			accountid,
			shippingid,
			street,
			city,
			stateprovince,
			postalcode,
			country)
			select
			ld.accountid,
			lsa.shippingid,
			lsa.street,
			lsa.city,
			lsa.stateprovince,
			lsa.postalcode,
			lsa.country
			from
			cms_leads_shipping_address_master lsa
			inner join
			cms_leads_master ld on ld.leadid=lsa.leadid
			where ld.leadid=$1`
		rows, err = r.db.DB().Query(sqlStatementasm1, lead.LeadId)
		fmt.Println("account Contact data is inserted into accounts_shipping_address_master")

		if err != nil {
			logger.Info(err.Error())
			return nil, err
			//return events.APIGatewayProxyResponse{500, headers, nil, err.Error(), false}, nil
		}
		defer rows.Close()

		res, _ := json.Marshal(rows)
		fmt.Println(res)
		return &model.AccountDetailsResponse{
			StatusCode: 200,
			Payload:    res,
		}, nil
		//return events.APIGatewayProxyResponse{200, headers, nil, string(res), false}, nil
	}

	res1, _ := json.Marshal("Success")
	fmt.Println(res1)
	//return events.APIGatewayProxyResponse{200, headers, nil, string(res1), false}, nil

	return &model.AccountDetailsResponse{
		StatusCode: 200,
		Payload:    res1,
	},nil
}
