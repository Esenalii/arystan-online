🔐 1. Авторизация және Тіркелу:
POST /register — жаңа қолданушы тіркеу

Пароль bcrypt арқылы хэштеледі

Егер роль берілмесе – "student" ретінде сақталады

POST /login — логин арқылы JWT токен алу

Жауапта token, name, role қайтарылады

🔒 2. Токен арқылы қорғау (middleware):
Authorization: Bearer <token> арқылы қорғауға дайын

JWTAuthMiddleware() — токенді тексереді, user_id, role контекстке салады

👥 3. User операциялары:
/api/v1/users — барлық қолданушыларды алу (мүмкін тек GET)

🎓 4. Student операциялары:
GET /api/v1/students — барлық студенттер

GET /api/v1/students/:id — нақты студентті алу

POST /api/v1/students — студент қосу

PUT /api/v1/students/:id — студентті жаңарту

DELETE /api/v1/students/:id — студентті өшіру

📘 5. Course операциялары:
GET /api/v1/courses — курс тізімін көру

POST /api/v1/courses — жаңа курс қосу