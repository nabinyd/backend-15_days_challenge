# ACESS 15-Day Learning Challenge

## Today, I started the **15-Day Backend Masterclass Challenge**


## Day 1: Database Design  

Focused on database design:  

- Designed a **Database Schema** using [dbdiagram.io](https://dbdiagram.io).  
- Created a structured **schema table** for a banking application backend.

---

## Day 2: Setup Local enviroment

 - Set up Docker and created a local development environment using PostgreSQL and SQLC Go.
 - Learned the basics of Docker, including commands to view images, containers, and logs.

---

## Day 3: Database Migration and Makefile Setup  
Today, I focused on improving the project structure and database management:  

- **Learned About:**  
  - Setting up database migrations using the `migrate` tool in Go.  
  - Creating a `Makefile` to streamline local development tasks and commands.  

- **Key Takeaway:**  
  Database migrations are essential for maintaining a consistent database schema throughout the development lifecycle. Using a `Makefile` simplifies repetitive tasks and improves project maintainability.  

---

## Day 4: Generating Go Code with SQLC  
The focus of today was on bridging SQL and Go seamlessly:  

- **Learned About:**  
  - Using **SQLC** to automatically generate Go code for SQL queries.  
  - Building a type-safe and efficient database interaction layer for the backend application.  

- **Key Takeaway:**  
  SQLC automates query-to-code generation, reducing errors and saving development time. It ensures consistency between the database and the application code, enhancing reliability.

---

## **Day 5: Unit Testing CRUD Functions**  
- Created unit test functions for CRUD operations.  
- Ensured comprehensive testing for Create, Read, Update, and Delete functionalities to improve code reliability.  
- Practiced mocking and validating edge cases to cover a wide range of scenarios.  

---

## **Day 6: Database Transactions and Deadlocks**  
- Learned about **database transactions** and their importance in ensuring data consistency and reliability.  
- Explored techniques to handle **deadlocks** effectively in concurrent database operations.  
- Improved understanding of multi-user database environments.  

---

## **Day 7: Handling and Avoiding Deadlocks**  
- Gained deeper insights into handling deadlocks in databases.  
- Started learning strategies to **avoid deadlocks** for building efficient and scalable backend systems.  

---

## Day 8: Adding Money Function & Isolation Levels  
- Added a new function to handle money additions in the backend service.  
- Studied **Database Isolation Levels** and their phenomena:  
  - **Dirty Reads**  
  - **Non-Repeatable Reads**  
  - **Phantom Reads**  

💡 Gained insights into how isolation levels impact database consistency and integrity.  

---

## Day 9: Practicing Read Uncommitted & Read Committed  
- Practiced **Read Uncommitted** and **Read Committed** isolation levels:  
  - **Read Uncommitted:** Allows dirty reads but provides minimal isolation.  
  - **Read Committed:** Prevents dirty reads by ensuring only committed data is read.  

💡 Improved understanding of isolation level trade-offs and their practical applications in database systems.  

---

## Day 10: Advanced Isolation Levels & GitHub Actions Setup  
- Learned about advanced transaction isolation levels:  
  - **Repeatable Read:** Ensures consistent reads by preventing non-repeatable reads.  
  - **Serializable:** Guarantees full isolation, eliminating phantom reads.  

- Configured **GitHub Actions** for CI/CD:  
  - Set up PostgreSQL in the CI environment.  
  - Integrated `go-migrate` for automated database migrations.  
  - Configured tests to run seamlessly in the GitHub Actions pipeline.  

💡 Gained practical experience in database management and automated testing workflows.  

---  

## Day 11: Setting Up Server and Routes with Pagination  

- Set up the **server** for the backend application.
- Created routes to:
  - **Get accounts**
  - **Create accounts**
  - **List accounts**
- Implemented **pagination** in the `list accounts` functionality to efficiently handle large datasets.

**Key Takeaway:**
- Learned the importance of proper server routing and the role of pagination in managing backend data effectively for better scalability and performance.

---

## **Day 12**  
- Loaded environment variables using Viper.  
- Learned to generate mocks using `mockgen` for more effective testing.  

---

## **Day 13**  
- Achieved 100% test coverage by creating comprehensive test cases.  
- Added custom validation logic for accounts to ensure data integrity.

---

## **Day 14**  
- Implemented a currency validator to check valid transactions.  
- Extended the database schema to include `User` data.  
- Learned to handle database errors effectively.

---

## **Day 15**  
- Implemented bcrypt-based password hashing for secure user data storage.  
- Created unit tests for the User API to ensure functionality and security.  

---

### **Completion and Future Goals**
Although this marks the end of the 15-day challenge, there's always more to learn in backend development. I'll continue exploring advanced topics and refining my skills.  



