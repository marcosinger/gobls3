# gobls3

## Description

It's tool for publish changes in my blog. It will automatically detect changed files and upload it to S3 bucket. 
For now it will update all changed files, but in future I'm planing make it more clever by adding caching.

## How it works

Tool will scan your blog folder from configuration file and recursively find files with last modified date. After gathering files, it will simply upload all of them to s3 bucket.

## Improvements

Of course it's not perfect and can be optimized/improvement:

- optimize folder scanning recursion algorithm

- cache the upload information, for minimize traffic in future

I'm planning to work on it in future.