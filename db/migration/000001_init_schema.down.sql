DROP TRIGGER "update_pay_out_modtime" ON "pay_out";

DROP TRIGGER "update_loan_modtime" ON "loan";

DROP TRIGGER "update_income_modtime" ON "income";

DROP TRIGGER "update_project_modtime" ON "project";

DROP FUNCTION "update_modified_column";

DROP TABLE IF EXISTS "pay_out";

DROP TABLE IF EXISTS "loan";

DROP TABLE IF EXISTS "income";

DROP TABLE IF EXISTS "project";

DROP TABLE IF EXISTS "users";

DROP EXTENSION IF EXISTS "uuid-ossp";
