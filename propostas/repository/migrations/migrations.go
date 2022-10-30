package migration

import (
	"database/sql"
	"fmt"
)

func CreateAttatchments(dbConn *sql.DB) {
	execSQL(dbConn, `CREATE TABLE IF NOT EXISTS public.attatchments (
		id serial4 NOT NULL,
		url varchar NOT NULL,
		user_id int8 NOT NULL
	);`, "Attatchments")
	return
}

func CreateMedia(dbConn *sql.DB) {
	execSQL(dbConn, `CREATE TABLE IF NOT EXISTS public.media (
		id serial4 NOT NULL,
		url varchar NOT NULL,
		user_id int8 NOT NULL,
		CONSTRAINT media_id PRIMARY KEY (id)
	);`, "Media")
	return

}

func CreateParllament(dbConn *sql.DB) {
	execSQL(dbConn, `CREATE TABLE IF NOT EXISTS public.parllament (
		id serial4 NOT NULL,
		email varchar(50) NOT NULL,
		template_name int8 NOT NULL
	);`, "Parllament")
	return
}

func CreateProposal(dbConn *sql.DB) {
	execSQL(dbConn, `CREATE TABLE IF NOT EXISTS public.proposal (
		id serial4 NOT NULL,
		title varchar(100) NOT NULL,
		pictures int8 NOT NULL,
		attachments int8 NOT NULL,
		description varchar(1000) NOT NULL,
		status bool NOT NULL,
		user_id int8 NOT NULL,
		CONSTRAINT proposal_id PRIMARY KEY (id)
	);`, "Proposal")
	return
}

func CreateUser(dbConn *sql.DB) {
	execSQL(dbConn, `CREATE TABLE IF NOT EXISTS public."user" (
		id serial4 NOT NULL,
		"name" varchar(50) NOT NULL,
		email varchar(50) NOT NULL,
		"password" varchar NOT NULL,
		picture int8 NOT NULL,
		newsletter bool NOT NULL,
		CONSTRAINT user_id PRIMARY KEY (id)
	);`, "User")
	return
}

func AlterAttatchments(dbConn *sql.DB) {
	execSQL(dbConn, `DO $$
	BEGIN
	
	BEGIN
		ALTER TABLE public.attatchments ADD CONSTRAINT media_fk FOREIGN KEY (user_id) REFERENCES public."user"(id);
	EXCEPTION
		WHEN duplicate_object THEN RAISE NOTICE 'Table constraint already exists';
	END;
	
  END $$;`, "Alter Attatchments")
	return
}

func AlterMedia(dbConn *sql.DB) {
	execSQL(dbConn, `DO $$
	BEGIN
	
	BEGIN
		ALTER TABLE public.attatchments ADD CONSTRAINT media_fk FOREIGN KEY (user_id) REFERENCES public."user"(id);
	EXCEPTION
		WHEN duplicate_object THEN RAISE NOTICE 'Table constraint already exists';
	END;
	
  END $$;`, "Alter Media")
	return
}

func AlterProposal(dbConn *sql.DB) {
	execSQL(dbConn, `DO $$
	BEGIN
	
	BEGIN
		ALTER TABLE public.proposal ADD CONSTRAINT proposal_attatchment FOREIGN KEY (attachments) REFERENCES public.media(id);
		ALTER TABLE public.proposal ADD CONSTRAINT proposal_fk FOREIGN KEY (user_id) REFERENCES public."user"(id);
		ALTER TABLE public.proposal ADD CONSTRAINT proposal_pictures FOREIGN KEY (pictures) REFERENCES public.media(id);
	EXCEPTION
		WHEN duplicate_object THEN RAISE NOTICE 'Table constraint already exists';
	END;
	
  END $$;`, "Alter Proposal")
	return
}

func AlterUser(dbConn *sql.DB) {
	execSQL(dbConn, `DO $$
		BEGIN
		
		BEGIN
			ALTER TABLE public."user" ADD CONSTRAINT user_fk FOREIGN KEY (picture) REFERENCES public.media(id);
		EXCEPTION
			WHEN duplicate_object THEN RAISE NOTICE 'Table constraint already exists';
		END;
		
	END $$;`, "Alter User")
	return
}

func execSQL(dbConn *sql.DB, query string, entityName string) {
	_, err := dbConn.Exec(query)
	if err != nil {
		panic(err)
		return
	}

	fmt.Println(entityName + " Migrated")
	return
}

func Exec(Conn *sql.DB) {
	CreateUser(Conn)
	CreateAttatchments(Conn)
	CreateMedia(Conn)
	CreateParllament(Conn)
	CreateProposal(Conn)
	AlterUser(Conn)
	AlterAttatchments(Conn)
	AlterMedia(Conn)
	AlterProposal(Conn)
}
