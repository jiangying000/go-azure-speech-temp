# before run:

```
export SPEECHSDK_ROOT="$HOME/speechsdk"
export LD_LIBRARY_PATH="$SPEECHSDK_ROOT/lib/x64:$LD_LIBRARY_PATH"
export CGO_CFLAGS="-I$SPEECHSDK_ROOT/include/c_api"
export CGO_LDFLAGS="-L$SPEECHSDK_ROOT/lib/x64 -lMicrosoft.CognitiveServices.Speech.core"

export SPEECH_KEY=YOU_KEY
export SPEECH_REGION=YOU_REGION
```