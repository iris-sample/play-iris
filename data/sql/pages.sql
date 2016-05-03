CREATE TABLE
    pages (
        id SERIAL PRIMARY KEY,
        title varchar(40) NOT NULL,
        content text NOT NULL,
        created TIMESTAMP
        );
INSERT INTO pages (title, content, created)
    VALUES ('GoLang Iris fan project', 'This project is for fun', CURRENT_TIMESTAMP);

INSERT INTO pages (title, content, created)
    VALUES (
        'What is Lorem Ipsum?',
        'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industrys standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.',
         CURRENT_TIMESTAMP
    );