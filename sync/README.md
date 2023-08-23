# sync

## Mutex vs Channel

Based on (the go wiki page extensively covering the topic)[https://github.com/golang/go/wiki/MutexOrChannel]

Paraphrasing:
Use channels when passing ownership of data
Use mutexes for managing state