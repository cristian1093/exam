CREATE OR REPLACE FUNCTION public.key_generator()
 RETURNS trigger
 LANGUAGE plpgsql
AS $function$
    BEGIN
		NEW.key = REPLACE(CONCAT(NEW.date, REPLACE(NEW.afilliation,'true',''), NEW.business_name, NEW.date_tx, NEW.import, NEW.account, REPLACE(NEW.ref,'true',''), NEW.acl_type, NEW.clarification_type, NEW.autorization, NEW.clarification_status, NEW.commitment_date, NEW.request_type, NEW.description, NEW.chargeback_code),' ','');

	    NEW.created_at = NOW();
	   RETURN NEW;  
    END;
$function$
;