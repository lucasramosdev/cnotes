BEGIN;

ALTER TABLE public.notes ALTER COLUMN title TYPE varchar(30) USING title::varchar(30);
ALTER TABLE public.annotations DROP CONSTRAINT annotations_clue_fk;
ALTER TABLE public.annotations DROP COLUMN IF EXISTS clue;
ALTER TABLE public.annotations ADD column "note" SERIAL NOT NULL;
ALTER TABLE public.annotations add CONSTRAINT annotations_note_fkey FOREIGN KEY (note) REFERENCES public.notes(id);
ALTER TABLE public.clues RENAME COLUMN value TO description;
ALTER TABLE public.clues RENAME TO keywords;

COMMIT;
