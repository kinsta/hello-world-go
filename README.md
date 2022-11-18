![Photo by Brett Ritchie on Unsplash](https://user-images.githubusercontent.com/2342458/202707249-359f3a1b-d61c-4d3b-a8b7-7a284f8f9842.png)

# Kinsta - Hello World - Go
An example of how to set your Go application up to enable deployment on Kinsta App Hosting services.

> Kinsta’s Application Hosting is a service to run your web apps and any databases side by side in a hassle-free environment, tailored for developer needs and ease of use. App Hosting is currently in an invite-only beta phase, sign up for a test account at [kinsta.com/application-hosting](https://kinsta.com/application-hosting/).

## Dependency Management

During the deployment process Kinsta will automatically install dependencies defined in your go.mod file.

## Web Server Setup

### Port
Kinsta automatically sets the `PORT` environment variable. You should **not** define it yourself and you should **not** hard-code it into the application. Use
`os.Getenv("PORT")` in your code when referring to the server port.

### Start Command
When deploying an application Kinsta will automatically create a web process.
