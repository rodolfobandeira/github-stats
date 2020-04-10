# github-stats
Gets statistics from Github

How to run?

```bash
GITHUB_TOKEN=<YOUR_TOKEN> GITHUB_USERNAME=<YOUR_USERNAME> GITHUB_REPOSITORY=<REPO> go run github-stats.go
```

How to build?

```bash
make
```

It will create the binary files inside `/bin` for the following platforms: Linux, ARM, Windows, MacOS


Current options:

```
---------------------------------------------------
Github Stats - Which statistics do you want to see?
---------------------------------------------------
1- Time to Merge Pull Requests
0- Exit
```
