CREATE TABLE lessons (
                         id SERIAL PRIMARY KEY,
                         created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
                         updated_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
                         deleted_at TIMESTAMP WITH TIME ZONE,
                         title TEXT NOT NULL,
                         content TEXT NOT NULL,
                         video TEXT,
                         course_id INT NOT NULL,
                         CONSTRAINT fk_course
                             FOREIGN KEY(course_id)
                                 REFERENCES courses(id)
                                 ON DELETE CASCADE
);