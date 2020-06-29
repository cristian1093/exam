CREATE OR REPLACE FUNCTION public.validate_afilliations()
 RETURNS trigger
 LANGUAGE plpgsql
AS $function$
    BEGIN
     IF NEW.date >= (select max(created_at::timestamp::date) + interval '15 day' from chargebacks where filename like '%INBURSA%') THEN 
   
   insert into chargebacks (chargeback_id,date, afilliation, business_name, date_tx, import, account, ref, acl_type, clarification_type, autorization, clarification_status, commitment_date, request_type, description, chargeback_code,filename) values (NEW.chargeback_id,NEW.date, REPLACE(NEW.afilliation,'true',''), NEW.business_name, NEW.date_tx, NEW.import, NEW.account, REPLACE(NEW.ref,'true',''), NEW.acl_type, NEW.clarification_type, NEW.autorization, NEW.clarification_status, NEW.commitment_date, NEW.request_type, NEW.description, NEW.chargeback_code,NEW.filename);
   
   ELSE  
    IF NEW.afilliation LIKE '%true%' and (CONCAT(REPLACE(NEW.afilliation,'true',''), NEW.business_name, NEW.date_tx, NEW.import, NEW.account, REPLACE(NEW.ref,'true',''), NEW.acl_type, NEW.clarification_type, NEW.autorization, NEW.clarification_status, NEW.commitment_date, NEW.request_type, NEW.description, NEW.chargeback_code)) NOT IN (select concat(REPLACE(afilliation,'true',''), business_name, date_tx, import, account, REPLACE(ref,'true',''), acl_type, clarification_type, autorization, clarification_status, commitment_date, request_type, description, chargeback_code) from chargebacks) THEN
        
     insert into chargebacks (chargeback_id,date, afilliation, business_name, date_tx, import, account, ref, acl_type, clarification_type, autorization, clarification_status, commitment_date, request_type, description, chargeback_code,filename) values (NEW.chargeback_id,NEW.date, REPLACE(NEW.afilliation,'true',''), NEW.business_name, NEW.date_tx, NEW.import, NEW.account, REPLACE(NEW.ref,'true',''), NEW.acl_type, NEW.clarification_type, NEW.autorization, NEW.clarification_status, NEW.commitment_date, NEW.request_type, NEW.description, NEW.chargeback_code,NEW.filename);
 
	 ELSE IF  NEW.afilliation NOT LIKE '%true%' THEN
	 
       insert into chargebacks_without_afilliation (date, afilliation, business_name, date_tx, import, account, ref, acl_type, clarification_type, autorization, clarification_status, commitment_date, request_type, description, chargeback_code,filename) values (NEW.date, REPLACE(NEW.afilliation,'true',''), NEW.business_name, NEW.date_tx, NEW.import, NEW.account, REPLACE(NEW.ref,'true',''), NEW.acl_type, NEW.clarification_type, NEW.autorization,NEW.clarification_status, NEW.commitment_date, NEW.request_type, NEW.description, NEW.chargeback_code,NEW.filename);
 
    ELSE IF  (CONCAT(REPLACE(NEW.afilliation,'true',''), NEW.business_name, NEW.date_tx, NEW.import, NEW.account, REPLACE(NEW.ref,'true',''), NEW.acl_type, NEW.clarification_type, NEW.autorization, NEW.clarification_status, NEW.commitment_date, NEW.request_type, NEW.description, NEW.chargeback_code)) IN (select concat(REPLACE(afilliation,'true',''), business_name, date_tx, import, account, REPLACE(ref,'true',''), acl_type, clarification_type, autorization, clarification_status, commitment_date, request_type, description, chargeback_code) from chargebacks)  THEN
	 
    insert into chargebacks_duplicates (date, afilliation, business_name, date_tx, import, account, ref, acl_type, clarification_type, autorization, clarification_status, commitment_date, request_type, description, chargeback_code,filename) values (NEW.date, REPLACE(NEW.afilliation,'true',''), NEW.business_name, NEW.date_tx, NEW.import, NEW.account, REPLACE(NEW.ref,'true',''), NEW.acl_type, NEW.clarification_type, NEW.autorization, NEW.clarification_status, NEW.commitment_date, NEW.request_type, NEW.description, NEW.chargeback_code,NEW.filename);
  	  END IF; 
   END IF;      
  END IF;
   END IF;
    	 RETURN NEW;   
    END;
$function$
;
