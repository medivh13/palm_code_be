# PALMCODE_HEADLESS_CMS

I use existing libs and tools :

 - Ozzo Validation, for input request validation
 - Godotenv, for env loader
 - jmoiron/sqlx for postgres driver
 - postgresql for DB
 - firebase storage for media's storage

## For setup after cloning/unzip the project:
 - cd palm_code_be
 - go mod tidy
 - add your firebase storage admin sdk file.json in to this root project
 - make changes in the .env file using your postgresql and your firebase storage bucket and firebase admin_sdk_path_to_your_root_project


## for db table :
 - in folder db, there are some .sql files with the create table command. I use postgresql for this case. you can run the command in your sql editor page

## to do a unit test :
 - i've made several unit testing in usecases/business layer
 - go to the each usecase package that you want to testing, then run a command "go test"
 - you can see the coverage testing in each usecase package by open the project with vscode, choose the testing file, Right click anywhere on the file display, then choose "Go:Toogle Test Coverage in Current Package"

## to run the project
 - after set the .env file with your database and firebase credential, then stay still in root directory, then do "go run main.go" in terminal

### API Endpoints
- **register**
```
curl --location 'http://localhost:8080/api/user/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email":"jody.almaida@gmail.com",
    "password": "12341234"
}'
```

- **login**
```
curl --location 'http://localhost:8080/api/user/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email":"jody.almaida@gmail.com",
    "password":"12341234"
    
}'
```

- **upload**
```
curl --location 'http://localhost:8080/api/upload/' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJleHAiOjE3MjEyOTY5NjJ9.9HgcxgZM_ATRnIOaye4zDFGvIvHzXbzBgDeH3-6hriE' \
--form 'file=@"/Users/jodyalmaida/Downloads/colage1.jpg"'
```

- **page Create**
```
curl --location 'http://localhost:8080/api/page' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJleHAiOjE3MjEyOTY5NjJ9.9HgcxgZM_ATRnIOaye4zDFGvIvHzXbzBgDeH3-6hriE' \
--header 'Content-Type: application/json' \
--data '{
    "title": "Example Page Title 1",
    "slug": "example-page-title 1",
    "banner_media": "https://storage.googleapis.com/download/storage/v1/b/palm-code-be-storage.appspot.com/o/colage1.jpg?generation=1721289809918562&alt=media",
    "content": "This is the content of the example page. It can be as long or as short as you need it to be.",
    "publication_date": "2024-07-18T12:00:00Z"
}
'
```
- **page Update**
```
curl --location --request PUT 'http://localhost:8080/api/page' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJleHAiOjE3MjEyOTY5NjJ9.9HgcxgZM_ATRnIOaye4zDFGvIvHzXbzBgDeH3-6hriE' \
--header 'Content-Type: application/json' \
--data '{
    "id": 4,
    "title": "Example Page Title2",
    "slug": "example-page-title2",
    "banner_media": "https://storage.googleapis.com/download/storage/v1/b/palm-code-be-storage.appspot.com/o/colage1.jpg?generation=1721289809918562&alt=media",
    "content": "This is the content of the example page. It can be as long or as short as you need it to be.",
    "publication_date": "2024-07-18T12:00:00Z"
}'
```

- **page Get**
```
curl --location 'http://localhost:8080/api/page?page=1&perPage=10' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJleHAiOjE3MjE0MDAxNTJ9.lywbBO-8PSdFzaOCSt81jKxAFascTr0PBukOVBFhZJY'
```

- **page Get By ID**
```
curl --location 'http://localhost:8080/api/page/1' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJleHAiOjE3MjEyOTY5NjJ9.9HgcxgZM_ATRnIOaye4zDFGvIvHzXbzBgDeH3-6hriE'
```

- **page Delete**
```
curl --location --request DELETE 'http://localhost:8080/api/page' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJleHAiOjE3MjEyOTY5NjJ9.9HgcxgZM_ATRnIOaye4zDFGvIvHzXbzBgDeH3-6hriE' \
--header 'Content-Type: application/json' \
--data '{
    "id": 4
}'
```

- **media Get**
```
curl --location 'http://localhost:8080/api/media?page=1&perPage=10' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJleHAiOjE3MjEyOTY5NjJ9.9HgcxgZM_ATRnIOaye4zDFGvIvHzXbzBgDeH3-6hriE'
```

- **media Get By ID**
```
curl --location 'http://localhost:8080/api/media/1' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJleHAiOjE3MjEyOTY5NjJ9.9HgcxgZM_ATRnIOaye4zDFGvIvHzXbzBgDeH3-6hriE'
```

- **team Create**
```
curl --location 'http://localhost:8080/api/team' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJleHAiOjE3MjEyOTY5NjJ9.9HgcxgZM_ATRnIOaye4zDFGvIvHzXbzBgDeH3-6hriE' \
--header 'Content-Type: application/json' \
--data '{
    "name": "John Doe 2",
    "role": "Developer",
    "bio": "Experienced software developer with a passion for creating scalable applications.",
    "profile_picture": "https://storage.googleapis.com/download/storage/v1/b/palm-code-be-storage.appspot.com/o/colage1.jpg?generation=1721289809918562&alt=media"
}
'
```

- **team Update**
```
curl --location --request PUT 'http://localhost:8080/api/team' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJleHAiOjE3MjEyOTY5NjJ9.9HgcxgZM_ATRnIOaye4zDFGvIvHzXbzBgDeH3-6hriE' \
--header 'Content-Type: application/json' \
--data '{
    "id":3,
    "name": "John Doe 3",
    "role": "Developer",
    "bio": "Experienced software developer with a passion for creating scalable applications.",
    "profile_picture": "https://storage.googleapis.com/download/storage/v1/b/palm-code-be-storage.appspot.com/o/colage1.jpg?generation=1721289809918562&alt=media"
}
'
```

- **team Get**
```
curl --location 'http://localhost:8080/api/team?page=1&perPage=10' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJleHAiOjE3MjEyOTY5NjJ9.9HgcxgZM_ATRnIOaye4zDFGvIvHzXbzBgDeH3-6hriE'
```

- **team Get By ID**
```
curl --location 'http://localhost:8080/api/team/2' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJleHAiOjE3MjEyOTY5NjJ9.9HgcxgZM_ATRnIOaye4zDFGvIvHzXbzBgDeH3-6hriE'
```

- **team Delete**
```
curl --location --request DELETE 'http://localhost:8080/api/team' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJleHAiOjE3MjEyOTY5NjJ9.9HgcxgZM_ATRnIOaye4zDFGvIvHzXbzBgDeH3-6hriE' \
--header 'Content-Type: application/json' \
--data '{
    "id": 2
}'
```

- **ping**
```
curl --location 'http://localhost:8080/ping'
```