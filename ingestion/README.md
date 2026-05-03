# Ingestion

# *** Ingestion Service ***

-> This is service which take care of uploading the song
-> master.wav / master.flac | This is the main file that is uploaded by artist
-> master.wav is the original recorded file
-> master.flac is the compressed file with same quality || compressed by finding the same pattern
-> these both are same just flac is space saver




-> now let us look at all the steps including in ingestion service
1. artist upload master file
2. client splits into chunks
3. Ingestion service receives chunks
4. chunks are temporarily stored
5. upload completion is verified
6. chunks are merged into master file
7. master file is validated 
8. master file is stored permanently 
9. encoding is trigerred
10. track is mark as ingested

-> now we will start exploring the steps 
1. Artists start upload
-> usually master.wav or master.flac
-> high quality
-> in future only from this file all variants are made

2. client splits into chunks
-> this is done by frontend
-> why in chunks 
-> suppose 10mb song
-> some error then resend whole song
-> but suppose 10 chunks -> then network error -> only that chunk you need to resend
-> each chunks has upload_id,  chunk_number, raw_bytes

3. Ingestion service receiving chunks
-> POST /ingestion/upload-chunk


4. Temporary chunks storage
-> chunks are stored in temporary storage
-> marked as incomplete
-> stored in local storage, object storage like s3
-> why temporary -> upload may be fail or may be cancelled

5. Upload completion verification
-> ingestion service checks -> are all chunks present, are continuous, size matches the actual size
-> this will reject or proceed

6. chunk merge
-> read chunks in order
-> append bytes
-> create one file -> master.flac

7. master file validation 
=> validiates 
-> file types
-> duration readable
-> no corruption 
-> audio headers valid
-> reject or continue

8. permanent storage of master
-> master file is stored in long term storage
-> object storage bucked, not cdn, not public 
-> source of truth 

9. Encoding

10. mark ingestion complete
-> track.status= ingested


-> ingestion service is done till step 8
-> at step 9 encoding service come into picture
-> ingestion service triggers encoding via messaging queue





*** Ingestion Service ***




POST /ingestion/init-upload

Response:
{
  upload_id,
  total_chunks,
  expected_size,
  file_type
}


Now there can be 2 types of uploads 
1. Server Upload
2. Direct Upload


1. *** Server Upload *** (Client -> Ingestion -> Storage)
-> now let us talk about server upload
-> Basically when upload starts
-> client will hit 'init' api then 'upload' api for each packet
-> your server | receive | validates | upload to storage

* Pros 
1. Full control
2. Easy frontend logic

* Cons
1. Scalability
-> 10k user start uploading song bakcend dies
2. Double bandwidth cost
-> Data travels twice
3. Latency Increases

* When to use 
1. Small apps
2. Internal tools

2. *** Direct Upload *** (Signed URL/ Multipart-Upload)
* Client → Storage (S3)
* Client → Backend (only metadata)

1. Client calls 
POST /ingestion/init-upload
2. Baclend returns
Signed URLS
3. Client Uploads directly to:
Storage
4. Client notifies backend

* Pros
1. Scalable
2. Cost efficient
3. Faster Upload

* Cons
1. More Complex client
2. Less Immediate Control
-> You can’t validate before upload
3. Security Complexity
-> Signed URL Expiry 
-> Permissions

This is the architecture we will follow now 
-> here we are combining both the approaches
1. init-upload → backend
2. upload chunks → S3 (direct)
3. complete-upload → backend
4. backend validates → triggers pipeline





Now let us look at our final design lifeCycle

1. Client will send Init request to Backend
-> api 
/init  Method Post

-> request
type initUploadRequest struct {
	ArtistID  string    `json:"artist_id"`
	FileName  string    `json:"file_name"`
	FileSize  int64     `json:"file_size"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

-> response
{
  "upload_id": "upl_123",
  "max_chunk_size": 5242880,
  "chunks": [
    { "part_number": 1, "url": "..." }
  ]
}


2. Client will put each chunk to Storage

3. Storage will send ETA tag to client

4. Client will call current chunk upload Completion
-> api
/mark-chunk-uploaded

-> request
{
  "upload_id": "upl_123",
  "chunk_number": 1,
  "etag": "abc123"
}

-> response
{}

5. Cycle repeat for each chunk

6. Upload Complete
-> api
/upload-complete
-> request 
{
  "upload_id" : "upl_123"
}
-> response 
{}