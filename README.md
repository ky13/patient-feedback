# patient-feedback

## Prompt

Please build a toy program to collect patient feedback after an appointment. Attached is some JSON with sample data about patients and their appointments which should drive the program. And below is a
sample script with what questions each patient should be asked.

We know we’ve given you a relatively small amount of time, and we want this to be fun. You just need to build something clean that functions and meets the requirements noted below.


## Sample Script

Program asks: “Hi [Patient First Name], on a scale of 1-10, would you recommend Dr [Doctor Last Name] to a friend or family member? 1 = Would not recommend, 10 = Would strongly recommend”

Assume patient responds with an integer between 1-10

Program asks: “Thank you. You were diagnosed with [Diagnosis]. Did Dr [Doctor Last Name] explain how to manage this diagnosis in a way you could understand?”

Record whatever the patient responds with, consider special treatment of Yes/No

Program asks: “We appreciate the feedback, one last question: how do you feel about being diagnosed with [Diagnosis]?”

Record whatever the patient responds with

Program ends with: “Thanks again! Here’s what we heard:”

Take some creative license with how patients are shown what the program recorded


## Requirements

1. Please make a basic architecture diagram
2. If you know any of the languages/technologies in our tech stack, please use them, i.e.
    1. If you know Go, code the back end in Go
    2. If you know React, build a React UI (otherwise a command line UI is fine)
    3. If you are familiar with cloud services (i.e. AWS), host the program in the cloud (otherwise running this locally is fine); AWS account access will be sent to you separately
    4. BUT you do NOT have to use Postgres or neo4j to store patient responses, you can if you want, but dumping patient responses to a JSON file or any database is fine
3. Please have at least one unit test or integration test that passes - we like TDD :)
4. Please prepare for code review, push your code up to GitHub or come prepared to show your code via an IDE, the code must be shared with us for future reference by the end of the interview
5. Please prepare to explain why you wrote your code the way you did - your thought process is as important as your code
6. Please prepare to explain how you worked through any technical problems you encountered
7. If you had to learn a new technology, please prepare to explain how you approached learning that new technology
8. Please prepare some thoughts on what you might have done differently, or what you would do next, if you were to keep working on this
9. Feel free to ask as many questions as you like, reach out if you ever get stuck, just reply all to this email
