# crawler-api
This is rest api for crawler project

# Running Locally with Docker Compose

## Prerequisites
- [Docker Desktop](https://www.docker.com/products/docker-desktop/) installed and running
- [Git](https://git-scm.com/) installed

## Steps

1. **Clone the repository:**
   ```sh
   git clone <your-repo-url>
   cd crawler-api
   ```

2. **Build and start the services:**
   ```sh
   docker-compose up --build
   ```
   This will build the Go application, start a MySQL database, and import the initial data from `db/mysql.sql`.

3. **Access the API:**
   - The API will be available at: [http://localhost:8080](http://localhost:8080)
   - MySQL will be available at port 13306 on your host (if you mapped it), or 3306 by default.

4. **Stopping the services:**
   Press `Ctrl+C` in the terminal, or run:
   ```sh
   docker-compose down
   ```

## Notes
- The backend will wait for the database to be ready before starting.
- Update `conf/config.yml` if you change database credentials or ports.
- For frontend development, ensure CORS is configured for your frontend's origin in `main.go`.
- User name: admin@skyell.com Password: 1234567

---
Replace `<your-repo-url>` with your actual GitHub repository URL.
