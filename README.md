# SimpleJobPortal

A Simple Job Portal API

## Installation

```bash
1. make sure you have go installed on your local machine
2. run 'git clone https://github.com/todyaja/SimpleJobPortal' in the desired folder
3. make sure to modify the env.go according to your local machine, especially the database port, host, username, password, and database name
4. run the sql codes on the sql folder through this order
    1. application_status.sql
    2. application_status_insert.sql
    3. user_type.sql
    4. user_type_insert.sql
    5. user_table.sql
    6. job_table.sql
    7. application_table.sql

5. run 'go run .' in terminal
```

## Endpoints

- /api/register
- /api/login

#Talent

- GET /api/talent/job
- POST /api/talent/apply-job
- GET /api/talent/my-application/{applicationId}
- PUT /api/employer/job/{jobId}
- DELETE /api/employer/job/{jobId}

#Employer

- POST /api/employer/job
- POST /api/employer/view-applicant
- PUT /api/employer/process-applicant
- POST /api/employer/withdraw-application

## Additional Notes

- There will be additional files added to this github, the additional files will be :
  1. ERD
  2. Postman Collection
  3. The Assesment Question
- Same email cant be registered twice, except its a different user type
- Operations on job and application (update, delete, view) can only be done by the creator
- Submitted application can't be updated

- Application Status
  1 = Ongoing
  2 = Interview
  3 = Accepted
  4 = Rejected

- User Type
  1 = Talent
  2 = Employer
