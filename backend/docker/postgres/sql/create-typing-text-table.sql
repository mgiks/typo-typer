CREATE TABLE IF NOT EXISTS typing_text (
    id serial PRIMARY KEY,
    content text NOT NULL,
    submitter varchar(100) NOT NULL,
    source text NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO
    typing_text (content, submitter, source)
VALUES
    (
        'Tom clapped his hands and cried: "Tom, Tom! Your guests are tired, and you had near forgotten! Come now, my merry friends, and Tom will refresh you! You shall clean grimy hands, and wash your weary faces; cast off your muddy cloaks and comb out your tangles!" He opened the door, and they followed him down a short passage and round a sharp turn. They came to a low room with a sloping roof (a penthouse, it seemed, built on the north end of the house). Its walls were of clean stone, but they were mostly covered with green hanging mats and yellow curtains. The floor was flagged, and strewn with fresh green rushes. There were four deep mattresses, each piled with white blankets, laid on the floor along one side. Against the opposite wall was a long bench laden with wide earthenware basins, and beside it stood brown ewers filled with water, some cold, some steaming hot. There were soft green slippers set ready beside each bed.',
        'mgik',
        'The Lord of The Rings'
    );
