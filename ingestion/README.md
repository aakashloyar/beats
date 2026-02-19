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

