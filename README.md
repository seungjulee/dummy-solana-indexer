# Overview
Here is a simple exercise into indexing Solana-like blockchain data. It indexes each account event into a SQLite database, logs the insertion, and calls the callback function. In case if there's a newer event with higher version comes in before the callback expires, the callback function gets cancelled.

## Prerequisite
1. sqlite

## Running it locally

Run indexer by `make run_indexer`


## Testing

Run indexer by `make test`

## Design decisions
A dedicated indexer program serves as if it's constantly ingesting a stream of data from a server. It simulates by using a coroutine channel with delays. It uses goroutines to fire and the callback within callback time. Using coroutine with context allows the go function to only print out either timeout result or the normal callback result.


## Future improvement
I would add more test cases with mocks about different scenarios of cancelling the callback. This can be done with checking if the expected calls were called after time has elapsed.

On SQL database, the schema could have a better indexing. If this was a production level, I would also use elasticsearch like realtime search tool to index the data as well.


## Example
```sh
 ✗ make run_indexer
go run ./cmd/indexer/main.go
2023-07-29T00:20:55.744+0900    INFO    initializing app with config    {"config": {"IndexerConfig":{"FakeAccountPath":"./input.json"},"SqliteConfig":{"SqlitePath":"./test.db"}}}
2023-07-29T00:20:55.746+0900    INFO    successfully connected to the db
2023-07-29T00:20:55.746+0900    INFO    migrate the schema tables
2023-07-29T00:20:55.747+0900    INFO    Scheduling periodic index
2023-07-29T00:20:55.747+0900    INFO    Successfully loaded account inputs
2023-07-29T00:20:56.499+0900    INFO    indexed account into db {"id": "6BhkGCMVMyrjEEkrASJcLxfAvoW43g6BubxjpeUyZFoz", "version": 5}
2023-07-29T00:20:57.000+0900    INFO    callback fires  {"id": "6BhkGCMVMyrjEEkrASJcLxfAvoW43g6BubxjpeUyZFoz", "version": 5}
2023-07-29T00:20:57.261+0900    INFO    indexed account into db {"id": "ZBTDeX9WPfT8P2KJ4mWhMXZccQPfpjuuq8DvEqQS7sB3", "version": 1}
2023-07-29T00:20:57.562+0900    INFO    callback fires  {"id": "ZBTDeX9WPfT8P2KJ4mWhMXZccQPfpjuuq8DvEqQS7sB3", "version": 1}
2023-07-29T00:20:58.093+0900    INFO    indexed account into db {"id": "U5kG5d2y4MJJEAxzzx8dCCLWJzzNs47eSo4Uo43eQVRg", "version": 1}
2023-07-29T00:20:58.193+0900    INFO    indexed account into db {"id": "N8F2F83EeFd7XLVeCMT4mhFiUgzSys3x5kWDpnmBKs9T", "version": 5}
2023-07-29T00:20:58.232+0900    INFO    indexed account into db {"id": "kbtvQBErkxRFePFB42p9V2fxRJc49LDuLDVGuc2zNfUF", "version": 1}
2023-07-29T00:20:58.713+0900    INFO    indexed account into db {"id": "cYcekbx6u4odny8L2ywkfc8R4wmGH7bX1GzNw7RagP6m", "version": 1}
2023-07-29T00:20:58.753+0900    INFO    indexed account into db {"id": "gkDF8QaRcyLCm4A6tsmHiQvqUZ7NqsvpXiZ8K8RQof9n", "version": 5}
2023-07-29T00:20:58.907+0900    INFO    indexed account into db {"id": "A72KnHfakikuhTGrbMg7dth9c9Sxjs3hypodQVDPNMR3", "version": 1}
2023-07-29T00:20:59.473+0900    INFO    indexed account into db {"id": "6BhkGCMVMyrjEEkrASJcLxfAvoW43g6BubxjpeUyZFoz", "version": 7}
2023-07-29T00:20:59.733+0900    INFO    callback fires  {"id": "kbtvQBErkxRFePFB42p9V2fxRJc49LDuLDVGuc2zNfUF", "version": 1}
2023-07-29T00:21:00.065+0900    INFO    indexed account into db {"id": "ZBTDeX9WPfT8P2KJ4mWhMXZccQPfpjuuq8DvEqQS7sB3", "version": 0}
2023-07-29T00:21:00.194+0900    INFO    indexed account into db {"id": "U5kG5d2y4MJJEAxzzx8dCCLWJzzNs47eSo4Uo43eQVRg", "version": 0}
2023-07-29T00:21:00.584+0900    INFO    indexed account into db {"id": "foXSdE2dWiip5Tw42QY94hcKSZRLZAaUmuSpYqsjANN8", "version": 7}
2023-07-29T00:21:00.733+0900    INFO    indexed account into db {"id": "2jymzXN3APMp3kjc44ESN35jFomF7hM59Ue8BMSv7pUG", "version": 0}
2023-07-29T00:21:00.805+0900    INFO    indexed account into db {"id": "cYcekbx6u4odny8L2ywkfc8R4wmGH7bX1GzNw7RagP6m", "version": 0}
2023-07-29T00:21:01.054+0900    INFO    callback fires  {"id": "gkDF8QaRcyLCm4A6tsmHiQvqUZ7NqsvpXiZ8K8RQof9n", "version": 5}
2023-07-29T00:21:01.799+0900    INFO    indexed account into db {"id": "gkDF8QaRcyLCm4A6tsmHiQvqUZ7NqsvpXiZ8K8RQof9n", "version": 7}
2023-07-29T00:21:02.185+0900    INFO    callback fires  {"id": "foXSdE2dWiip5Tw42QY94hcKSZRLZAaUmuSpYqsjANN8", "version": 7}
2023-07-29T00:21:02.506+0900    INFO    indexed account into db {"id": "A72KnHfakikuhTGrbMg7dth9c9Sxjs3hypodQVDPNMR3", "version": 0}
2023-07-29T00:21:02.563+0900    INFO    indexed account into db {"id": "6BhkGCMVMyrjEEkrASJcLxfAvoW43g6BubxjpeUyZFoz", "version": 9}
2023-07-29T00:21:02.563+0900    INFO    callback cancelled      {"id": "6BhkGCMVMyrjEEkrASJcLxfAvoW43g6BubxjpeUyZFoz", "version": 7}
2023-07-29T00:21:02.960+0900    INFO    indexed account into db {"id": "ZBTDeX9WPfT8P2KJ4mWhMXZccQPfpjuuq8DvEqQS7sB3", "version": 3}
2023-07-29T00:21:02.960+0900    INFO    callback cancelled      {"id": "ZBTDeX9WPfT8P2KJ4mWhMXZccQPfpjuuq8DvEqQS7sB3", "version": 0}
2023-07-29T00:21:03.215+0900    INFO    callback fires  {"id": "cYcekbx6u4odny8L2ywkfc8R4wmGH7bX1GzNw7RagP6m", "version": 1}
2023-07-29T00:21:03.290+0900    INFO    indexed account into db {"id": "xo5mm1Lsmy5b2c4jFyKR4SbSkCxXUKBabPYT6TLzUJZK", "version": 3}
2023-07-29T00:21:03.494+0900    INFO    callback fires  {"id": "U5kG5d2y4MJJEAxzzx8dCCLWJzzNs47eSo4Uo43eQVRg", "version": 1}
2023-07-29T00:21:03.734+0900    INFO    callback fires  {"id": "2jymzXN3APMp3kjc44ESN35jFomF7hM59Ue8BMSv7pUG", "version": 0}
2023-07-29T00:21:03.940+0900    INFO    indexed account into db {"id": "foXSdE2dWiip5Tw42QY94hcKSZRLZAaUmuSpYqsjANN8", "version": 9}
2023-07-29T00:21:04.866+0900    INFO    indexed account into db {"id": "2jymzXN3APMp3kjc44ESN35jFomF7hM59Ue8BMSv7pUG", "version": 3}
2023-07-29T00:21:05.139+0900    INFO    indexed account into db {"id": "hhpGbCqzxJDCCHEDFXXD3b8XUbTRUygDpc36qQZdy7pL", "version": 3}
2023-07-29T00:21:05.741+0900    INFO    callback fires  {"id": "foXSdE2dWiip5Tw42QY94hcKSZRLZAaUmuSpYqsjANN8", "version": 9}
2023-07-29T00:21:05.861+0900    INFO    callback fires  {"id": "ZBTDeX9WPfT8P2KJ4mWhMXZccQPfpjuuq8DvEqQS7sB3", "version": 3}
2023-07-29T00:21:05.984+0900    INFO    indexed account into db {"id": "gkDF8QaRcyLCm4A6tsmHiQvqUZ7NqsvpXiZ8K8RQof9n", "version": 9}
2023-07-29T00:21:05.984+0900    INFO    callback cancelled      {"id": "gkDF8QaRcyLCm4A6tsmHiQvqUZ7NqsvpXiZ8K8RQof9n", "version": 7}
2023-07-29T00:21:06.185+0900    INFO    callback fires  {"id": "gkDF8QaRcyLCm4A6tsmHiQvqUZ7NqsvpXiZ8K8RQof9n", "version": 9}
2023-07-29T00:21:06.924+0900    INFO    indexed account into db {"id": "A72KnHfakikuhTGrbMg7dth9c9Sxjs3hypodQVDPNMR3", "version": 3}
2023-07-29T00:21:06.924+0900    INFO    callback cancelled      {"id": "A72KnHfakikuhTGrbMg7dth9c9Sxjs3hypodQVDPNMR3", "version": 0}
2023-07-29T00:21:07.008+0900    INFO    callback fires  {"id": "A72KnHfakikuhTGrbMg7dth9c9Sxjs3hypodQVDPNMR3", "version": 1}
2023-07-29T00:21:07.240+0900    INFO    indexed account into db {"id": "6BhkGCMVMyrjEEkrASJcLxfAvoW43g6BubxjpeUyZFoz", "version": 11}
2023-07-29T00:21:07.240+0900    INFO    callback cancelled      {"id": "6BhkGCMVMyrjEEkrASJcLxfAvoW43g6BubxjpeUyZFoz", "version": 9}
2023-07-29T00:21:07.295+0900    INFO    callback fires  {"id": "U5kG5d2y4MJJEAxzzx8dCCLWJzzNs47eSo4Uo43eQVRg", "version": 0}
2023-07-29T00:21:07.883+0900    INFO    indexed account into db {"id": "ZBTDeX9WPfT8P2KJ4mWhMXZccQPfpjuuq8DvEqQS7sB3", "version": 2}
2023-07-29T00:21:07.994+0900    INFO    callback fires  {"id": "N8F2F83EeFd7XLVeCMT4mhFiUgzSys3x5kWDpnmBKs9T", "version": 5}
2023-07-29T00:21:08.134+0900    INFO    indexed account into db {"id": "U5kG5d2y4MJJEAxzzx8dCCLWJzzNs47eSo4Uo43eQVRg", "version": 2}
2023-07-29T00:21:08.493+0900    INFO    indexed account into db {"id": "N8F2F83EeFd7XLVeCMT4mhFiUgzSys3x5kWDpnmBKs9T", "version": 11}
2023-07-29T00:21:08.906+0900    INFO    callback fires  {"id": "cYcekbx6u4odny8L2ywkfc8R4wmGH7bX1GzNw7RagP6m", "version": 0}
2023-07-29T00:21:09.266+0900    INFO    callback fires  {"id": "2jymzXN3APMp3kjc44ESN35jFomF7hM59Ue8BMSv7pUG", "version": 3}
2023-07-29T00:21:09.304+0900    INFO    indexed account into db {"id": "2jymzXN3APMp3kjc44ESN35jFomF7hM59Ue8BMSv7pUG", "version": 2}
2023-07-29T00:21:09.872+0900    INFO    indexed account into db {"id": "cYcekbx6u4odny8L2ywkfc8R4wmGH7bX1GzNw7RagP6m", "version": 2}
2023-07-29T00:21:09.925+0900    INFO    callback fires  {"id": "A72KnHfakikuhTGrbMg7dth9c9Sxjs3hypodQVDPNMR3", "version": 3}
2023-07-29T00:21:10.121+0900    INFO    indexed account into db {"id": "gkDF8QaRcyLCm4A6tsmHiQvqUZ7NqsvpXiZ8K8RQof9n", "version": 11}
2023-07-29T00:21:10.254+0900    INFO    indexed account into db {"id": "A72KnHfakikuhTGrbMg7dth9c9Sxjs3hypodQVDPNMR3", "version": 2}
2023-07-29T00:21:11.097+0900    INFO    indexed account into db {"id": "6BhkGCMVMyrjEEkrASJcLxfAvoW43g6BubxjpeUyZFoz", "version": 13}
2023-07-29T00:21:11.097+0900    INFO    callback cancelled      {"id": "6BhkGCMVMyrjEEkrASJcLxfAvoW43g6BubxjpeUyZFoz", "version": 11}
2023-07-29T00:21:11.391+0900    INFO    callback fires  {"id": "xo5mm1Lsmy5b2c4jFyKR4SbSkCxXUKBabPYT6TLzUJZK", "version": 3}
2023-07-29T00:21:11.605+0900    INFO    callback fires  {"id": "2jymzXN3APMp3kjc44ESN35jFomF7hM59Ue8BMSv7pUG", "version": 2}
2023-07-29T00:21:11.916+0900    INFO    indexed account into db {"id": "ZBTDeX9WPfT8P2KJ4mWhMXZccQPfpjuuq8DvEqQS7sB3", "version": 5}
2023-07-29T00:21:11.916+0900    INFO    callback cancelled      {"id": "ZBTDeX9WPfT8P2KJ4mWhMXZccQPfpjuuq8DvEqQS7sB3", "version": 2}
2023-07-29T00:21:12.560+0900    INFO    indexed account into db {"id": "xo5mm1Lsmy5b2c4jFyKR4SbSkCxXUKBabPYT6TLzUJZK", "version": 5}
2023-07-29T00:21:13.221+0900    INFO    indexed account into db {"id": "N8F2F83EeFd7XLVeCMT4mhFiUgzSys3x5kWDpnmBKs9T", "version": 13}
2023-07-29T00:21:13.221+0900    INFO    callback cancelled      {"id": "N8F2F83EeFd7XLVeCMT4mhFiUgzSys3x5kWDpnmBKs9T", "version": 11}
2023-07-29T00:21:13.856+0900    INFO    indexed account into db {"id": "kbtvQBErkxRFePFB42p9V2fxRJc49LDuLDVGuc2zNfUF", "version": 5}
2023-07-29T00:21:14.162+0900    INFO    callback fires  {"id": "xo5mm1Lsmy5b2c4jFyKR4SbSkCxXUKBabPYT6TLzUJZK", "version": 5}
2023-07-29T00:21:14.441+0900    INFO    callback fires  {"id": "hhpGbCqzxJDCCHEDFXXD3b8XUbTRUygDpc36qQZdy7pL", "version": 3}
2023-07-29T00:21:14.506+0900    INFO    indexed account into db {"id": "cYcekbx6u4odny8L2ywkfc8R4wmGH7bX1GzNw7RagP6m", "version": 5}
2023-07-29T00:21:14.506+0900    INFO    callback cancelled      {"id": "cYcekbx6u4odny8L2ywkfc8R4wmGH7bX1GzNw7RagP6m", "version": 2}
2023-07-29T00:21:14.587+0900    INFO    indexed account into db {"id": "gkDF8QaRcyLCm4A6tsmHiQvqUZ7NqsvpXiZ8K8RQof9n", "version": 13}
2023-07-29T00:21:14.587+0900    INFO    callback cancelled      {"id": "gkDF8QaRcyLCm4A6tsmHiQvqUZ7NqsvpXiZ8K8RQof9n", "version": 11}
2023-07-29T00:21:14.945+0900    INFO    indexed account into db {"id": "A72KnHfakikuhTGrbMg7dth9c9Sxjs3hypodQVDPNMR3", "version": 5}
2023-07-29T00:21:14.946+0900    INFO    callback cancelled      {"id": "A72KnHfakikuhTGrbMg7dth9c9Sxjs3hypodQVDPNMR3", "version": 2}
2023-07-29T00:21:15.488+0900    INFO    callback fires  {"id": "gkDF8QaRcyLCm4A6tsmHiQvqUZ7NqsvpXiZ8K8RQof9n", "version": 13}
2023-07-29T00:21:15.534+0900    INFO    indexed account into db {"id": "6BhkGCMVMyrjEEkrASJcLxfAvoW43g6BubxjpeUyZFoz", "version": 15}
2023-07-29T00:21:15.534+0900    INFO    callback cancelled      {"id": "6BhkGCMVMyrjEEkrASJcLxfAvoW43g6BubxjpeUyZFoz", "version": 13}
2023-07-29T00:21:15.747+0900    INFO    callback fires  {"id": "A72KnHfakikuhTGrbMg7dth9c9Sxjs3hypodQVDPNMR3", "version": 5}
2023-07-29T00:21:16.135+0900    INFO    callback fires  {"id": "U5kG5d2y4MJJEAxzzx8dCCLWJzzNs47eSo4Uo43eQVRg", "version": 2}
2023-07-29T00:21:16.141+0900    INFO    indexed account into db {"id": "ZBTDeX9WPfT8P2KJ4mWhMXZccQPfpjuuq8DvEqQS7sB3", "version": 4}
2023-07-29T00:21:16.445+0900    INFO    indexed account into db {"id": "xo5mm1Lsmy5b2c4jFyKR4SbSkCxXUKBabPYT6TLzUJZK", "version": 4}
2023-07-29T00:21:16.891+0900    INFO    indexed account into db {"id": "N8F2F83EeFd7XLVeCMT4mhFiUgzSys3x5kWDpnmBKs9T", "version": 15}
2023-07-29T00:21:16.891+0900    INFO    callback cancelled      {"id": "N8F2F83EeFd7XLVeCMT4mhFiUgzSys3x5kWDpnmBKs9T", "version": 13}
2023-07-29T00:21:17.500+0900    INFO    indexed account into db {"id": "2jymzXN3APMp3kjc44ESN35jFomF7hM59Ue8BMSv7pUG", "version": 4}
2023-07-29T00:21:18.169+0900    INFO    indexed account into db {"id": "hhpGbCqzxJDCCHEDFXXD3b8XUbTRUygDpc36qQZdy7pL", "version": 4}
2023-07-29T00:21:18.417+0900    INFO    callback fires  {"id": "ZBTDeX9WPfT8P2KJ4mWhMXZccQPfpjuuq8DvEqQS7sB3", "version": 5}
2023-07-29T00:21:18.446+0900    INFO    callback fires  {"id": "xo5mm1Lsmy5b2c4jFyKR4SbSkCxXUKBabPYT6TLzUJZK", "version": 4}
2023-07-29T00:21:18.487+0900    INFO    indexed account into db {"id": "gkDF8QaRcyLCm4A6tsmHiQvqUZ7NqsvpXiZ8K8RQof9n", "version": 15}
2023-07-29T00:21:18.635+0900    INFO    callback fires  {"id": "6BhkGCMVMyrjEEkrASJcLxfAvoW43g6BubxjpeUyZFoz", "version": 15}
2023-07-29T00:21:18.942+0900    INFO    callback fires  {"id": "ZBTDeX9WPfT8P2KJ4mWhMXZccQPfpjuuq8DvEqQS7sB3", "version": 4}
2023-07-29T00:21:19.025+0900    INFO    indexed account into db {"id": "A72KnHfakikuhTGrbMg7dth9c9Sxjs3hypodQVDPNMR3", "version": 4}
2023-07-29T00:21:20.032+0900    INFO    indexed account into db {"id": "6BhkGCMVMyrjEEkrASJcLxfAvoW43g6BubxjpeUyZFoz", "version": 17}
2023-07-29T00:21:20.871+0900    INFO    indexed account into db {"id": "ZBTDeX9WPfT8P2KJ4mWhMXZccQPfpjuuq8DvEqQS7sB3", "version": 7}
2023-07-29T00:21:21.001+0900    INFO    callback fires  {"id": "2jymzXN3APMp3kjc44ESN35jFomF7hM59Ue8BMSv7pUG", "version": 4}
2023-07-29T00:21:21.584+0900    INFO    indexed account into db {"id": "U5kG5d2y4MJJEAxzzx8dCCLWJzzNs47eSo4Uo43eQVRg", "version": 7}
2023-07-29T00:21:21.658+0900    INFO    callback fires  {"id": "kbtvQBErkxRFePFB42p9V2fxRJc49LDuLDVGuc2zNfUF", "version": 5}
2023-07-29T00:21:21.788+0900    INFO    callback fires  {"id": "gkDF8QaRcyLCm4A6tsmHiQvqUZ7NqsvpXiZ8K8RQof9n", "version": 15}
2023-07-29T00:21:21.907+0900    INFO    callback fires  {"id": "cYcekbx6u4odny8L2ywkfc8R4wmGH7bX1GzNw7RagP6m", "version": 5}
2023-07-29T00:21:22.022+0900    INFO    indexed account into db {"id": "foXSdE2dWiip5Tw42QY94hcKSZRLZAaUmuSpYqsjANN8", "version": 17}
2023-07-29T00:21:22.426+0900    INFO    callback fires  {"id": "A72KnHfakikuhTGrbMg7dth9c9Sxjs3hypodQVDPNMR3", "version": 4}
2023-07-29T00:21:22.844+0900    INFO    indexed account into db {"id": "2jymzXN3APMp3kjc44ESN35jFomF7hM59Ue8BMSv7pUG", "version": 7}
2023-07-29T00:21:22.939+0900    INFO    indexed account into db {"id": "hhpGbCqzxJDCCHEDFXXD3b8XUbTRUygDpc36qQZdy7pL", "version": 7}
2023-07-29T00:21:22.939+0900    INFO    callback cancelled      {"id": "hhpGbCqzxJDCCHEDFXXD3b8XUbTRUygDpc36qQZdy7pL", "version": 4}
2023-07-29T00:21:22.985+0900    INFO    callback fires  {"id": "U5kG5d2y4MJJEAxzzx8dCCLWJzzNs47eSo4Uo43eQVRg", "version": 7}
2023-07-29T00:21:23.164+0900    INFO    indexed account into db {"id": "gkDF8QaRcyLCm4A6tsmHiQvqUZ7NqsvpXiZ8K8RQof9n", "version": 17}
2023-07-29T00:21:23.697+0900    INFO    indexed account into db {"id": "A72KnHfakikuhTGrbMg7dth9c9Sxjs3hypodQVDPNMR3", "version": 7}
2023-07-29T00:21:24.054+0900    INFO    indexed account into db {"id": "6BhkGCMVMyrjEEkrASJcLxfAvoW43g6BubxjpeUyZFoz", "version": 19}
2023-07-29T00:21:24.054+0900    INFO    callback cancelled      {"id": "6BhkGCMVMyrjEEkrASJcLxfAvoW43g6BubxjpeUyZFoz", "version": 17}
2023-07-29T00:21:24.054+0900    INFO    callback fires  {"id": "6BhkGCMVMyrjEEkrASJcLxfAvoW43g6BubxjpeUyZFoz", "version": 19}
2023-07-29T00:21:24.148+0900    INFO    indexed account into db {"id": "ZBTDeX9WPfT8P2KJ4mWhMXZccQPfpjuuq8DvEqQS7sB3", "version": 6}
2023-07-29T00:21:24.903+0900    INFO    indexed account into db {"id": "xo5mm1Lsmy5b2c4jFyKR4SbSkCxXUKBabPYT6TLzUJZK", "version": 6}
2023-07-29T00:21:25.428+0900    INFO    indexed account into db {"id": "N8F2F83EeFd7XLVeCMT4mhFiUgzSys3x5kWDpnmBKs9T", "version": 19}
2023-07-29T00:21:25.429+0900    INFO    callback cancelled      {"id": "N8F2F83EeFd7XLVeCMT4mhFiUgzSys3x5kWDpnmBKs9T", "version": 15}
2023-07-29T00:21:26.426+0900    INFO    indexed account into db {"id": "2jymzXN3APMp3kjc44ESN35jFomF7hM59Ue8BMSv7pUG", "version": 6}
2023-07-29T00:21:26.550+0900    INFO    indexed account into db {"id": "hhpGbCqzxJDCCHEDFXXD3b8XUbTRUygDpc36qQZdy7pL", "version": 6}
2023-07-29T00:21:26.842+0900    INFO    callback fires  {"id": "hhpGbCqzxJDCCHEDFXXD3b8XUbTRUygDpc36qQZdy7pL", "version": 7}
2023-07-29T00:21:27.250+0900    INFO    indexed account into db {"id": "gkDF8QaRcyLCm4A6tsmHiQvqUZ7NqsvpXiZ8K8RQof9n", "version": 19}
2023-07-29T00:21:27.250+0900    INFO    callback cancelled      {"id": "gkDF8QaRcyLCm4A6tsmHiQvqUZ7NqsvpXiZ8K8RQof9n", "version": 17}
2023-07-29T00:21:27.304+0900    INFO    callback fires  {"id": "xo5mm1Lsmy5b2c4jFyKR4SbSkCxXUKBabPYT6TLzUJZK", "version": 6}
2023-07-29T00:21:28.250+0900    INFO    indexed account into db {"id": "A72KnHfakikuhTGrbMg7dth9c9Sxjs3hypodQVDPNMR3", "version": 6}
2023-07-29T00:21:28.372+0900    INFO    callback fires  {"id": "ZBTDeX9WPfT8P2KJ4mWhMXZccQPfpjuuq8DvEqQS7sB3", "version": 7}
2023-07-29T00:21:28.830+0900    INFO    callback fires  {"id": "N8F2F83EeFd7XLVeCMT4mhFiUgzSys3x5kWDpnmBKs9T", "version": 19}
2023-07-29T00:21:29.049+0900    INFO    indexed account into db {"id": "6BhkGCMVMyrjEEkrASJcLxfAvoW43g6BubxjpeUyZFoz", "version": 21}
2023-07-29T00:21:29.423+0900    INFO    callback fires  {"id": "foXSdE2dWiip5Tw42QY94hcKSZRLZAaUmuSpYqsjANN8", "version": 17}
2023-07-29T00:21:29.484+0900    INFO    indexed account into db {"id": "ZBTDeX9WPfT8P2KJ4mWhMXZccQPfpjuuq8DvEqQS7sB3", "version": 9}
2023-07-29T00:21:29.484+0900    INFO    callback cancelled      {"id": "ZBTDeX9WPfT8P2KJ4mWhMXZccQPfpjuuq8DvEqQS7sB3", "version": 6}
2023-07-29T00:21:29.949+0900    INFO    indexed account into db {"id": "U5kG5d2y4MJJEAxzzx8dCCLWJzzNs47eSo4Uo43eQVRg", "version": 9}
2023-07-29T00:21:30.526+0900    INFO    indexed account into db {"id": "N8F2F83EeFd7XLVeCMT4mhFiUgzSys3x5kWDpnmBKs9T", "version": 21}
2023-07-29T00:21:30.985+0900    INFO    callback fires  {"id": "ZBTDeX9WPfT8P2KJ4mWhMXZccQPfpjuuq8DvEqQS7sB3", "version": 9}
2023-07-29T00:21:31.122+0900    INFO    indexed account into db {"id": "kbtvQBErkxRFePFB42p9V2fxRJc49LDuLDVGuc2zNfUF", "version": 9}
2023-07-29T00:21:31.695+0900    INFO    indexed account into db {"id": "cYcekbx6u4odny8L2ywkfc8R4wmGH7bX1GzNw7RagP6m", "version": 9}
2023-07-29T00:21:31.908+0900    INFO    indexed account into db {"id": "gkDF8QaRcyLCm4A6tsmHiQvqUZ7NqsvpXiZ8K8RQof9n", "version": 21}
2023-07-29T00:21:31.908+0900    INFO    callback cancelled      {"id": "gkDF8QaRcyLCm4A6tsmHiQvqUZ7NqsvpXiZ8K8RQof9n", "version": 19}
2023-07-29T00:21:32.245+0900    INFO    callback fires  {"id": "2jymzXN3APMp3kjc44ESN35jFomF7hM59Ue8BMSv7pUG", "version": 7}
2023-07-29T00:21:32.527+0900    INFO    callback fires  {"id": "2jymzXN3APMp3kjc44ESN35jFomF7hM59Ue8BMSv7pUG", "version": 6}
2023-07-29T00:21:32.582+0900    INFO    indexed account into db {"id": "A72KnHfakikuhTGrbMg7dth9c9Sxjs3hypodQVDPNMR3", "version": 9}
2023-07-29T00:21:32.582+0900    INFO    callback cancelled      {"id": "A72KnHfakikuhTGrbMg7dth9c9Sxjs3hypodQVDPNMR3", "version": 6}
2023-07-29T00:21:32.807+0900    INFO    indexed account into db {"id": "6BhkGCMVMyrjEEkrASJcLxfAvoW43g6BubxjpeUyZFoz", "version": 23}
2023-07-29T00:21:32.807+0900    INFO    callback cancelled      {"id": "6BhkGCMVMyrjEEkrASJcLxfAvoW43g6BubxjpeUyZFoz", "version": 21}
2023-07-29T00:21:32.988+0900    INFO    indexed account into db {"id": "ZBTDeX9WPfT8P2KJ4mWhMXZccQPfpjuuq8DvEqQS7sB3", "version": 8}
2023-07-29T00:21:33.151+0900    INFO    callback fires  {"id": "hhpGbCqzxJDCCHEDFXXD3b8XUbTRUygDpc36qQZdy7pL", "version": 6}
2023-07-29T00:21:33.298+0900    INFO    callback fires  {"id": "A72KnHfakikuhTGrbMg7dth9c9Sxjs3hypodQVDPNMR3", "version": 7}
2023-07-29T00:21:33.312+0900    INFO    indexed account into db {"id": "U5kG5d2y4MJJEAxzzx8dCCLWJzzNs47eSo4Uo43eQVRg", "version": 8}
2023-07-29T00:21:33.796+0900    INFO    callback fires  {"id": "cYcekbx6u4odny8L2ywkfc8R4wmGH7bX1GzNw7RagP6m", "version": 9}
2023-07-29T00:21:33.858+0900    INFO    indexed account into db {"id": "N8F2F83EeFd7XLVeCMT4mhFiUgzSys3x5kWDpnmBKs9T", "version": 23}
2023-07-29T00:21:33.858+0900    INFO    callback cancelled      {"id": "N8F2F83EeFd7XLVeCMT4mhFiUgzSys3x5kWDpnmBKs9T", "version": 21}
2023-07-29T00:21:34.146+0900    INFO    indexed account into db {"id": "kbtvQBErkxRFePFB42p9V2fxRJc49LDuLDVGuc2zNfUF", "version": 8}
2023-07-29T00:21:34.259+0900    INFO    callback fires  {"id": "N8F2F83EeFd7XLVeCMT4mhFiUgzSys3x5kWDpnmBKs9T", "version": 23}
2023-07-29T00:21:34.270+0900    INFO    indexed account into db {"id": "cYcekbx6u4odny8L2ywkfc8R4wmGH7bX1GzNw7RagP6m", "version": 8}
2023-07-29T00:21:34.550+0900    INFO    callback fires  {"id": "U5kG5d2y4MJJEAxzzx8dCCLWJzzNs47eSo4Uo43eQVRg", "version": 9}
2023-07-29T00:21:34.797+0900    INFO    indexed account into db {"id": "gkDF8QaRcyLCm4A6tsmHiQvqUZ7NqsvpXiZ8K8RQof9n", "version": 23}
2023-07-29T00:21:34.798+0900    INFO    callback cancelled      {"id": "gkDF8QaRcyLCm4A6tsmHiQvqUZ7NqsvpXiZ8K8RQof9n", "version": 21}
2023-07-29T00:21:35.015+0900    INFO    callback fires  {"id": "U5kG5d2y4MJJEAxzzx8dCCLWJzzNs47eSo4Uo43eQVRg", "version": 8}
2023-07-29T00:21:35.113+0900    INFO    indexed account into db {"id": "A72KnHfakikuhTGrbMg7dth9c9Sxjs3hypodQVDPNMR3", "version": 8}
2023-07-29T00:21:35.389+0900    INFO    callback fires  {"id": "ZBTDeX9WPfT8P2KJ4mWhMXZccQPfpjuuq8DvEqQS7sB3", "version": 8}
2023-07-29T00:21:35.986+0900    INFO    indexed account into db {"id": "6BhkGCMVMyrjEEkrASJcLxfAvoW43g6BubxjpeUyZFoz", "version": 25}
2023-07-29T00:21:35.986+0900    INFO    callback cancelled      {"id": "6BhkGCMVMyrjEEkrASJcLxfAvoW43g6BubxjpeUyZFoz", "version": 23}
2023-07-29T00:21:36.371+0900    INFO    callback fires  {"id": "cYcekbx6u4odny8L2ywkfc8R4wmGH7bX1GzNw7RagP6m", "version": 8}
2023-07-29T00:21:36.796+0900    INFO    indexed account into db {"id": "ZBTDeX9WPfT8P2KJ4mWhMXZccQPfpjuuq8DvEqQS7sB3", "version": 11}
2023-07-29T00:21:37.299+0900    INFO    callback fires  {"id": "gkDF8QaRcyLCm4A6tsmHiQvqUZ7NqsvpXiZ8K8RQof9n", "version": 23}
2023-07-29T00:21:37.365+0900    INFO    indexed account into db {"id": "xo5mm1Lsmy5b2c4jFyKR4SbSkCxXUKBabPYT6TLzUJZK", "version": 11}
2023-07-29T00:21:37.712+0900    INFO    indexed account into db {"id": "foXSdE2dWiip5Tw42QY94hcKSZRLZAaUmuSpYqsjANN8", "version": 25}
2023-07-29T00:21:37.784+0900    INFO    callback fires  {"id": "A72KnHfakikuhTGrbMg7dth9c9Sxjs3hypodQVDPNMR3", "version": 9}
2023-07-29T00:21:37.969+0900    INFO    indexed account into db {"id": "kbtvQBErkxRFePFB42p9V2fxRJc49LDuLDVGuc2zNfUF", "version": 11}
2023-07-29T00:21:37.969+0900    INFO    callback cancelled      {"id": "kbtvQBErkxRFePFB42p9V2fxRJc49LDuLDVGuc2zNfUF", "version": 8}
2023-07-29T00:21:38.751+0900    INFO    indexed account into db {"id": "cYcekbx6u4odny8L2ywkfc8R4wmGH7bX1GzNw7RagP6m", "version": 11}
2023-07-29T00:21:38.874+0900    INFO    indexed account into db {"id": "gkDF8QaRcyLCm4A6tsmHiQvqUZ7NqsvpXiZ8K8RQof9n", "version": 25}
2023-07-29T00:21:39.106+0900    INFO    indexed account into db {"id": "A72KnHfakikuhTGrbMg7dth9c9Sxjs3hypodQVDPNMR3", "version": 11}
2023-07-29T00:21:39.106+0900    INFO    callback cancelled      {"id": "A72KnHfakikuhTGrbMg7dth9c9Sxjs3hypodQVDPNMR3", "version": 8}
2023-07-29T00:21:39.696+0900    INFO    indexed account into db {"id": "6BhkGCMVMyrjEEkrASJcLxfAvoW43g6BubxjpeUyZFoz", "version": 27}
2023-07-29T00:21:39.696+0900    INFO    callback cancelled      {"id": "6BhkGCMVMyrjEEkrASJcLxfAvoW43g6BubxjpeUyZFoz", "version": 25}
2023-07-29T00:21:40.689+0900    INFO    indexed account into db {"id": "ZBTDeX9WPfT8P2KJ4mWhMXZccQPfpjuuq8DvEqQS7sB3", "version": 10}
2023-07-29T00:21:41.023+0900    INFO    callback fires  {"id": "kbtvQBErkxRFePFB42p9V2fxRJc49LDuLDVGuc2zNfUF", "version": 9}
2023-07-29T00:21:41.206+0900    INFO    indexed account into db {"id": "xo5mm1Lsmy5b2c4jFyKR4SbSkCxXUKBabPYT6TLzUJZK", "version": 10}
2023-07-29T00:21:41.397+0900    INFO    callback fires  {"id": "ZBTDeX9WPfT8P2KJ4mWhMXZccQPfpjuuq8DvEqQS7sB3", "version": 11}
2023-07-29T00:21:42.161+0900    INFO    indexed account into db {"id": "foXSdE2dWiip5Tw42QY94hcKSZRLZAaUmuSpYqsjANN8", "version": 27}
2023-07-29T00:21:42.161+0900    INFO    callback cancelled      {"id": "foXSdE2dWiip5Tw42QY94hcKSZRLZAaUmuSpYqsjANN8", "version": 25}
2023-07-29T00:21:42.512+0900    INFO    indexed account into db {"id": "kbtvQBErkxRFePFB42p9V2fxRJc49LDuLDVGuc2zNfUF", "version": 10}
2023-07-29T00:21:42.552+0900    INFO    callback fires  {"id": "cYcekbx6u4odny8L2ywkfc8R4wmGH7bX1GzNw7RagP6m", "version": 11}
2023-07-29T00:21:42.694+0900    INFO    indexed account into db {"id": "hhpGbCqzxJDCCHEDFXXD3b8XUbTRUygDpc36qQZdy7pL", "version": 10}
2023-07-29T00:21:42.770+0900    INFO    callback fires  {"id": "kbtvQBErkxRFePFB42p9V2fxRJc49LDuLDVGuc2zNfUF", "version": 11}
2023-07-29T00:21:43.288+0900    INFO    indexed account into db {"id": "gkDF8QaRcyLCm4A6tsmHiQvqUZ7NqsvpXiZ8K8RQof9n", "version": 27}
2023-07-29T00:21:43.288+0900    INFO    callback cancelled      {"id": "gkDF8QaRcyLCm4A6tsmHiQvqUZ7NqsvpXiZ8K8RQof9n", "version": 25}
2023-07-29T00:21:43.691+0900    INFO    callback fires  {"id": "ZBTDeX9WPfT8P2KJ4mWhMXZccQPfpjuuq8DvEqQS7sB3", "version": 10}
2023-07-29T00:21:44.064+0900    INFO    indexed account into db {"id": "A72KnHfakikuhTGrbMg7dth9c9Sxjs3hypodQVDPNMR3", "version": 10}
2023-07-29T00:21:44.332+0900    INFO    indexed account into db {"id": "6BhkGCMVMyrjEEkrASJcLxfAvoW43g6BubxjpeUyZFoz", "version": 29}
2023-07-29T00:21:44.332+0900    INFO    callback cancelled      {"id": "6BhkGCMVMyrjEEkrASJcLxfAvoW43g6BubxjpeUyZFoz", "version": 27}
2023-07-29T00:21:44.980+0900    INFO    indexed account into db {"id": "ZBTDeX9WPfT8P2KJ4mWhMXZccQPfpjuuq8DvEqQS7sB3", "version": 0}
2023-07-29T00:21:45.333+0900    INFO    callback fires  {"id": "6BhkGCMVMyrjEEkrASJcLxfAvoW43g6BubxjpeUyZFoz", "version": 29}
2023-07-29T00:21:45.666+0900    INFO    indexed account into db {"id": "xo5mm1Lsmy5b2c4jFyKR4SbSkCxXUKBabPYT6TLzUJZK", "version": 0}
2023-07-29T00:21:46.208+0900    INFO    callback fires  {"id": "A72KnHfakikuhTGrbMg7dth9c9Sxjs3hypodQVDPNMR3", "version": 11}
2023-07-29T00:21:46.300+0900    INFO    indexed account into db {"id": "N8F2F83EeFd7XLVeCMT4mhFiUgzSys3x5kWDpnmBKs9T", "version": 29}
2023-07-29T00:21:46.486+0900    INFO    indexed account into db {"id": "2jymzXN3APMp3kjc44ESN35jFomF7hM59Ue8BMSv7pUG", "version": 0}
2023-07-29T00:21:46.695+0900    INFO    indexed account into db {"id": "cYcekbx6u4odny8L2ywkfc8R4wmGH7bX1GzNw7RagP6m", "version": 0}
2023-07-29T00:21:46.766+0900    INFO    callback fires  {"id": "xo5mm1Lsmy5b2c4jFyKR4SbSkCxXUKBabPYT6TLzUJZK", "version": 11}
2023-07-29T00:21:47.092+0900    INFO    indexed account into db {"id": "gkDF8QaRcyLCm4A6tsmHiQvqUZ7NqsvpXiZ8K8RQof9n", "version": 29}
2023-07-29T00:21:47.092+0900    INFO    callback cancelled      {"id": "gkDF8QaRcyLCm4A6tsmHiQvqUZ7NqsvpXiZ8K8RQof9n", "version": 27}
2023-07-29T00:21:47.213+0900    INFO    callback fires  {"id": "kbtvQBErkxRFePFB42p9V2fxRJc49LDuLDVGuc2zNfUF", "version": 10}
2023-07-29T00:21:47.767+0900    INFO    callback fires  {"id": "xo5mm1Lsmy5b2c4jFyKR4SbSkCxXUKBabPYT6TLzUJZK", "version": 0}
2023-07-29T00:21:47.808+0900    INFO    indexed account into db {"id": "A72KnHfakikuhTGrbMg7dth9c9Sxjs3hypodQVDPNMR3", "version": 0}
2023-07-29T00:21:47.893+0900    INFO    callback fires  {"id": "gkDF8QaRcyLCm4A6tsmHiQvqUZ7NqsvpXiZ8K8RQof9n", "version": 29}
2023-07-29T00:21:48.112+0900    INFO    callback fires  {"id": "A72KnHfakikuhTGrbMg7dth9c9Sxjs3hypodQVDPNMR3", "version": 0}
2023-07-29T00:21:48.127+0900    INFO    indexed account into db {"id": "6BhkGCMVMyrjEEkrASJcLxfAvoW43g6BubxjpeUyZFoz", "version": 31}
2023-07-29T00:21:48.522+0900    INFO    indexed account into db {"id": "ZBTDeX9WPfT8P2KJ4mWhMXZccQPfpjuuq8DvEqQS7sB3", "version": 12}
2023-07-29T00:21:48.523+0900    INFO    callback cancelled      {"id": "ZBTDeX9WPfT8P2KJ4mWhMXZccQPfpjuuq8DvEqQS7sB3", "version": 0}
2023-07-29T00:21:48.637+0900    INFO    indexed account into db {"id": "U5kG5d2y4MJJEAxzzx8dCCLWJzzNs47eSo4Uo43eQVRg", "version": 12}
2023-07-29T00:21:48.786+0900    INFO    indexed account into db {"id": "foXSdE2dWiip5Tw42QY94hcKSZRLZAaUmuSpYqsjANN8", "version": 31}
2023-07-29T00:21:48.786+0900    INFO    callback cancelled      {"id": "foXSdE2dWiip5Tw42QY94hcKSZRLZAaUmuSpYqsjANN8", "version": 27}
2023-07-29T00:21:49.287+0900    INFO    callback fires  {"id": "2jymzXN3APMp3kjc44ESN35jFomF7hM59Ue8BMSv7pUG", "version": 0}
2023-07-29T00:21:49.524+0900    INFO    callback fires  {"id": "ZBTDeX9WPfT8P2KJ4mWhMXZccQPfpjuuq8DvEqQS7sB3", "version": 12}
2023-07-29T00:21:49.595+0900    INFO    callback fires  {"id": "hhpGbCqzxJDCCHEDFXXD3b8XUbTRUygDpc36qQZdy7pL", "version": 10}
2023-07-29T00:21:49.775+0900    INFO    indexed account into db {"id": "2jymzXN3APMp3kjc44ESN35jFomF7hM59Ue8BMSv7pUG", "version": 12}
2023-07-29T00:21:49.897+0900    INFO    callback fires  {"id": "cYcekbx6u4odny8L2ywkfc8R4wmGH7bX1GzNw7RagP6m", "version": 0}
2023-07-29T00:21:50.727+0900    INFO    indexed account into db {"id": "hhpGbCqzxJDCCHEDFXXD3b8XUbTRUygDpc36qQZdy7pL", "version": 12}
2023-07-29T00:21:51.106+0900    INFO    callback fires  {"id": "xo5mm1Lsmy5b2c4jFyKR4SbSkCxXUKBabPYT6TLzUJZK", "version": 10}
2023-07-29T00:21:51.291+0900    INFO    indexed account into db {"id": "gkDF8QaRcyLCm4A6tsmHiQvqUZ7NqsvpXiZ8K8RQof9n", "version": 31}
2023-07-29T00:21:51.401+0900    INFO    callback fires  {"id": "N8F2F83EeFd7XLVeCMT4mhFiUgzSys3x5kWDpnmBKs9T", "version": 29}
2023-07-29T00:21:51.896+0900    INFO    indexed account into db {"id": "A72KnHfakikuhTGrbMg7dth9c9Sxjs3hypodQVDPNMR3", "version": 12}
2023-07-29T00:21:52.465+0900    INFO    callback fires  {"id": "A72KnHfakikuhTGrbMg7dth9c9Sxjs3hypodQVDPNMR3", "version": 10}
2023-07-29T00:21:52.693+0900    INFO    indexed account into db {"id": "6BhkGCMVMyrjEEkrASJcLxfAvoW43g6BubxjpeUyZFoz", "version": 33}
2023-07-29T00:21:52.693+0900    INFO    callback fires  {"id": "6BhkGCMVMyrjEEkrASJcLxfAvoW43g6BubxjpeUyZFoz", "version": 33}
2023-07-29T00:21:52.693+0900    INFO    callback cancelled      {"id": "6BhkGCMVMyrjEEkrASJcLxfAvoW43g6BubxjpeUyZFoz", "version": 31}
2023-07-29T00:21:52.938+0900    INFO    callback fires  {"id": "U5kG5d2y4MJJEAxzzx8dCCLWJzzNs47eSo4Uo43eQVRg", "version": 12}
2023-07-29T00:21:53.176+0900    INFO    callback fires  {"id": "2jymzXN3APMp3kjc44ESN35jFomF7hM59Ue8BMSv7pUG", "version": 12}
2023-07-29T00:21:53.407+0900    INFO    indexed account into db {"id": "N8F2F83EeFd7XLVeCMT4mhFiUgzSys3x5kWDpnmBKs9T", "version": 33}
2023-07-29T00:21:53.621+0900    INFO    indexed account into db {"id": "gkDF8QaRcyLCm4A6tsmHiQvqUZ7NqsvpXiZ8K8RQof9n", "version": 33}
2023-07-29T00:21:53.621+0900    INFO    callback cancelled      {"id": "gkDF8QaRcyLCm4A6tsmHiQvqUZ7NqsvpXiZ8K8RQof9n", "version": 31}
2023-07-29T00:21:54.080+0900    INFO    indexed account into db {"id": "6BhkGCMVMyrjEEkrASJcLxfAvoW43g6BubxjpeUyZFoz", "version": 35}
2023-07-29T00:21:54.945+0900    INFO    indexed account into db {"id": "foXSdE2dWiip5Tw42QY94hcKSZRLZAaUmuSpYqsjANN8", "version": 35}
2023-07-29T00:21:54.946+0900    INFO    callback cancelled      {"id": "foXSdE2dWiip5Tw42QY94hcKSZRLZAaUmuSpYqsjANN8", "version": 31}
2023-07-29T00:21:55.619+0900    INFO    indexed account into db {"id": "gkDF8QaRcyLCm4A6tsmHiQvqUZ7NqsvpXiZ8K8RQof9n", "version": 35}
2023-07-29T00:21:55.619+0900    INFO    callback cancelled      {"id": "gkDF8QaRcyLCm4A6tsmHiQvqUZ7NqsvpXiZ8K8RQof9n", "version": 33}
2023-07-29T00:21:56.538+0900    INFO    indexed account into db {"id": "6BhkGCMVMyrjEEkrASJcLxfAvoW43g6BubxjpeUyZFoz", "version": 37}
2023-07-29T00:21:56.539+0900    INFO    callback cancelled      {"id": "6BhkGCMVMyrjEEkrASJcLxfAvoW43g6BubxjpeUyZFoz", "version": 35}
2023-07-29T00:21:56.998+0900    INFO    callback fires  {"id": "A72KnHfakikuhTGrbMg7dth9c9Sxjs3hypodQVDPNMR3", "version": 12}
2023-07-29T00:21:57.439+0900    INFO    indexed account into db {"id": "foXSdE2dWiip5Tw42QY94hcKSZRLZAaUmuSpYqsjANN8", "version": 37}
2023-07-29T00:21:57.439+0900    INFO    callback cancelled      {"id": "foXSdE2dWiip5Tw42QY94hcKSZRLZAaUmuSpYqsjANN8", "version": 35}
2023-07-29T00:21:58.293+0900    INFO    indexed account into db {"id": "gkDF8QaRcyLCm4A6tsmHiQvqUZ7NqsvpXiZ8K8RQof9n", "version": 37}
2023-07-29T00:21:58.293+0900    INFO    callback cancelled      {"id": "gkDF8QaRcyLCm4A6tsmHiQvqUZ7NqsvpXiZ8K8RQof9n", "version": 35}
2023-07-29T00:21:58.571+0900    INFO    indexed account into db {"id": "6BhkGCMVMyrjEEkrASJcLxfAvoW43g6BubxjpeUyZFoz", "version": 39}
2023-07-29T00:21:58.571+0900    INFO    callback cancelled      {"id": "6BhkGCMVMyrjEEkrASJcLxfAvoW43g6BubxjpeUyZFoz", "version": 37}
2023-07-29T00:21:59.339+0900    INFO    indexed account into db {"id": "foXSdE2dWiip5Tw42QY94hcKSZRLZAaUmuSpYqsjANN8", "version": 39}
2023-07-29T00:21:59.339+0900    INFO    callback cancelled      {"id": "foXSdE2dWiip5Tw42QY94hcKSZRLZAaUmuSpYqsjANN8", "version": 37}
2023-07-29T00:21:59.621+0900    INFO    indexed account into db {"id": "gkDF8QaRcyLCm4A6tsmHiQvqUZ7NqsvpXiZ8K8RQof9n", "version": 39}
2023-07-29T00:21:59.621+0900    INFO    callback cancelled      {"id": "gkDF8QaRcyLCm4A6tsmHiQvqUZ7NqsvpXiZ8K8RQof9n", "version": 37}
2023-07-29T00:22:00.028+0900    INFO    callback fires  {"id": "hhpGbCqzxJDCCHEDFXXD3b8XUbTRUygDpc36qQZdy7pL", "version": 12}
2023-07-29T00:22:01.372+0900    INFO    callback fires  {"id": "6BhkGCMVMyrjEEkrASJcLxfAvoW43g6BubxjpeUyZFoz", "version": 39}
2023-07-29T00:22:02.109+0900    INFO    callback fires  {"id": "N8F2F83EeFd7XLVeCMT4mhFiUgzSys3x5kWDpnmBKs9T", "version": 33}
2023-07-29T00:22:03.721+0900    INFO    callback fires  {"id": "gkDF8QaRcyLCm4A6tsmHiQvqUZ7NqsvpXiZ8K8RQof9n", "version": 39}
2023-07-29T00:22:04.440+0900    INFO    callback fires  {"id": "foXSdE2dWiip5Tw42QY94hcKSZRLZAaUmuSpYqsjANN8", "version": 39}
2023-07-29T00:22:09.622+0900    INFO    finished ingesting events and processing event callbacks
2023-07-29T00:22:09.623+0900    INFO    the highest token-value accounts by AccountType {"account_id": "cYcekbx6u4odny8L2ywkfc8R4wmGH7bX1GzNw7RagP6m", "token": 960}
```
