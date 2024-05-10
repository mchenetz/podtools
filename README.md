# podtools
Description: Podtools allows you to create files needed for a podcast or other show. Currently supported functions:

1. Convert video to audio
2. Transcribe audio file
3. Use ChatGPT to react to a query on a text file. This is good for things like summarizing, crating a description, etc...



Usage:
The initial config sets a yaml file in your home directory that will save your apikey for futre use. You can also save defaults like input and output filenames and default directories. CLI flags will override any defaults when used.

# Set OpenAI API key


```
podtools config -a [apikey]
```




# Convert a file from Video to Audio
```
podtools convert -i [file] -o [file]
```
# Transcribe a file using OpenAI Whisper

```
podtools transcribe -i [file]
```
# Use OpenAI to return a query based on a text file
```
podtools -i [file] -q "Provide a youtube description from the attached podcast episode"
```









