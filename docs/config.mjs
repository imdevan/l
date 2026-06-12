const stage = process.env.NODE_ENV || "dev"
const isProduction = stage === "production"

export default {
  url: isProduction ? "https://devan.gg" : "http://localhost:4321",
  basePath:  isProduction ? "/l" : "/",
  github: "https://github.com/imdevan/l/",
  githubDocs: "https://github.com/imdevan/l/",
  title: "l",
  description: "an ls replacement",
}
