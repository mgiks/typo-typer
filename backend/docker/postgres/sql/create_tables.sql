CREATE TYPE text_category AS ENUM ('long', 'regular', 'short');

CREATE TABLE IF NOT EXISTS typing_text (
    id serial PRIMARY KEY,
    text text NOT NULL,
    category text_category NOT NULL
);

INSERT INTO
    typing_text (text, category)
VALUES
    (
        'Tom clapped his hands and cried: "Tom, Tom! Your guests are tired, and you had near forgotten! Come now, my merry friends, and Tom will refresh you! You shall clean grimy hands, and wash your weary faces; cast off your muddy cloaks and comb out your tangles!" He opened the door, and they followed him down a short passage and round a sharp turn. They came to a low room with a sloping roof (a penthouse, it seemed, built on the north end of the house). Its walls were of clean stone, but they were mostly covered with green hanging mats and yellow curtains. The floor was flagged, and strewn with fresh green rushes. There were four deep mattresses, each piled with white blankets, laid on the floor along one side. Against the opposite wall was a long bench laden with wide earthenware basins, and beside it stood brown ewers filled with water, some cold, some steaming hot. There were soft green slippers set ready beside each bed.',
        'long'
    ),
    (
        'Atoms of radioactive elements can split. According to Albert Einstein, mass and energy are interchangeable under certain circumstances. When atoms split, the process is called nuclear fission. In this case, a small amount of mass is converted into energy. Thus the energy released cannot do much damage. However, several subatomic particles called neutrons are also emitted during this process. Each neutron will hit a radioactive element releasing more neutrons in the process. This causes a chain reaction and creates a large amount of energy. This energy is converted into heat which expands uncontrollably causing an explosion. Hence, atoms do not literally explode. They generate energy that can cause explosions.',
        'regular'
    ),
    (
        'Here''s an account of how a man really lost his balance.',
        'short'
    );
