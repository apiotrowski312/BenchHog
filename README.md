# BenchHog

### To do:
* [x] Simple but informative metrics:
  * [x] Status codes (percentage of ok/error)
  * [x] average
  * [x] median
  * [x] percentage of times (like 50% under 1000seconds)
* [x] Basic Unit tests
* [x] Flags:
  * [x] URL
  * [x] number of total requests
  * [x] number of concurent requests
* [x] Support multiple endpoints within one bench test
  * [ ] Add flag with multiple tactics (like random, one-by-one etc)
* [ ] Redo request if get non 2xx request
  * [ ] add flag to disable `redo`
* [ ] Post
* [ ] Put