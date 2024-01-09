# Optimizing integration tests using template1 in PostgreSQL
This repository demonstrates a highly efficient method for managing test databases in Go using PostgreSQL's template feature. By leveraging this feature, we can significantly reduce the time required to run integration tests, especially in a continuous integration (CI) environment.

# Overview
The primary goal of this project is to showcase a pattern that reduces the overhead typically associated with setting up and tearing down test databases. Instead of running costly migrations for each test, we utilize a pre-migrated template database, allowing us to quickly clone its structure. This approach can cut down test execution times by up to 95%, turning minutes into mere seconds.

# Detailed explanation
For a detailed explanation of how the template feature works and the benefits it brings to CI pipelines, check out my Medium article: [How I Reduced Integration Test Execution Times by 95% Using a Secret PostgreSQL Gem.](https://medium.com/@mmacovv/how-i-reduced-integration-test-execution-times-by-95-using-a-secret-postgresql-gem-ebfaaeb96ed3)

# Getting started
To get started, you'll need to have [Docker](https://www.docker.com/) installed on your machine. Once you have Docker installed, you can run the following command to start a PostgreSQL instance:
```bash
docker compose up --build -d
```

Furthermore you also need to install Go 1.21 or higher. Once you have Go installed, you can run the following command to run the tests:
```bash
go test -bench=.
```

# Contribution
Contributions to this project are welcome! If you have improvements or suggestions, please open an issue or submit a pull request. For questions or discussions, feel free to reach out to me directly on [LinkedIn](https://www.linkedin.com/in/mohamedmacow/).

