### A base Go template app that has OAuth2 for login and JWT token for authentication.

#### Steps
1. Create an app on Facebook. Note appId and app secret.
2. Fulfil env variables listed in `src/init.go`.
3. Install dependencies by running `make deps`.
4. Run `make watch` to build and start server locally.
5. Visit `http://localhost:8000/`.
6. Any changes made in code will be taken care by `make watch`.
7. Go build.
