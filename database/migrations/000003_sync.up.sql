BEGIN;

ALTER TABLE public.notes ALTER COLUMN title TYPE text USING title::text;
ALTER TABLE public.keywords RENAME TO clues;
ALTER TABLE public.annotations ADD clue SERIAL NOT NULL;
ALTER TABLE public.annotations ADD CONSTRAINT annotations_clue_fk FOREIGN KEY (id) REFERENCES public.clues(id);
ALTER TABLE public.annotations DROP CONSTRAINT annotations_note_fkey;
ALTER TABLE public.annotations DROP COLUMN note;

ALTER TABLE public.clues RENAME COLUMN description TO value;

COMMIT;