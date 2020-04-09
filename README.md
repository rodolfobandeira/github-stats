# github-stats
Gets statistics from Github

How to run?

```bash
GITHUB_TOKEN=<YOUR_TOKEN> GITHUB_USERNAME=<YOUR_USERNAME> GITHUB_REPOSITORY=<REPO> go run github-stats.go
```


Current options:

```
---------------------------------------------------
Github Stats - Which statistics do you want to see?
---------------------------------------------------
1- Time to Merge Pull Requests
0- Exit
```

Output:

```
Title: TITLE PR 1
Merged in: 3.000394 days

Title: TITLE PR 2
Merged in: 2.000394 days


Total Pull Requests: 2
Average Time to Merge: 2.5 days
```
