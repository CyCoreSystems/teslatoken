# Tesla Token

In order to not hand out your username and password to third parties, it is a
reasonably common practice to instead give them a temporal API token which
expires.  This isn't great, because there is no permission control, but until
Tesla actually implements a proper system for this, this is what we have.

I had been using someone's Python script for this, but it quit working... so
this is just a quick hack to do the same in Go.

