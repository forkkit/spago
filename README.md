spaGO
=====
[![Go Report Card](https://goreportcard.com/badge/github.com/nlpodyssey/spago)](https://goreportcard.com/report/github.com/nlpodyssey/spago)

spaGO is a machine learning lightweight open-source library written in Go designed to support relevant neural network architectures in natural language processing tasks.

Are you looking for a highly optimized, scalable, battle-tested production-ready machine-learning/NLP framework? Are you also a Python lover and enjoy manipulating tensors? If yes, you won't find much to your satisfaction here. [PyTorch](https://pytorch.org/) plus the wonders of the friends of [Hugging Face](https://github.com/huggingface) is the answer you seek!

If instead you prefer statically typed, compiled programming language, and a **simpler yet well-structured** machine-learning framework almost ready to use is what you need, then you are in the right place!
The idea is that you could have written spaGO. Most of it, from the computational graph to the LSTM is straightforward Go code :)

### Note

**This README is under construction, you might consider coming back here in a few days, or go straight to the code if you're curious!**

Installation
=====
Make sure you have [Go 1.14](https://golang.org/dl/) installed on your computer first. The package can be installed using *go get* as follows:

```console
go get -u https://github.com/nlpodyssey/spago
```

spaGO is compatible with [go modules](https://blog.golang.org/using-go-modules).

Why spaGO?
=====

I've been writing more or less the same software for almost 20 years. I guess it's my way of learning a new language. Now it's Go's turn, and spaGO is the result of a few days of pure fun!

Let me explain a little further. It's not precisely the very same software I've been writing now for 20 years: I've been working in the NLP for this long, experimenting with different approaches and techniques, and therefore software of the same field. 
I've always felt satisfied to limit the use of third-party dependencies, trying to implement firsthand the features that interest me most. 
For instance, nowadays, you're nobody if you don't master machine learning. Seriously, NLP's state-of-the-art results have all been achieved with this approach! So, I took the opportunity to speed up my understanding of the underlying deep learning algorithms implementing them almost from scratch in straightforward Go code.
I'm aware that [reinventing the wheel](https://en.wikipedia.org/wiki/Reinventing_the_wheel#Related_phrases) is an anti-pattern; nevertheless, I wanted to build something with my own concepts in my own (italian) style: that's the way I learn best, and it could be your best chance to understand what's going on under the hood of the artificial intelligence :)

When I begin programming in a new language, I know almost anything about it. I often combine the techniques I have acquired by writing in other languages and other paradigms, so some choices may not be the most idiomatic... but who cares, right? 

It's with this approach that I jumped on Go and created spaGo: a work in progress, (hopefully) understandable, easy to use library for machine learning and natural language processing.

Disclaimer
=====

**Please note that I can only do development in my free time** (which is very limited: I am a [#onewheeler](https://twitter.com/hashtag/onewheel), I have a wonderful wife, a [Happy](https://github.com/nlpodyssey/spago/blob/master/assets/happy.jpg) dog, I play the piano and the guitar, and last but not least I'm actively engaged in my [daily job](https://www.exop-group.com/en/)), so no promises are made regarding response time, feature implementations or bug fixes.
If you want spaGo to become something more than just a hobby project of me, I greatly appreciate any bug reports and contributions, which can be made by filing an issue or making a pull request through the github page. Thanks!

Licensing
=====

spaGO is licensed under a BSD-style license. See [LICENSE](https://github.com/nlpodyssey/spago/blob/master/LICENSE) for the full license text.