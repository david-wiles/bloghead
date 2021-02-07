# BLOGHEAD

A basic CLI interface for creating a static website, with a focus on blog creation.

## The problem

I want to create a blog of my own using a static site generator, but I don't want to learn how to use an existing 
platform such as Hugo or Eleventy (as good as those projects are).

This project was equally motivated by a desire to build something and to create an easy to use interface for creating a 
static website. My first attempt at this can be found at [github.com/david-wiles/cmd-cv](https://github.com/david-wiles/cmd-cv).
While this serves my purposes well, I wanted to be able to easily create a new article from a partial HTML file. My 
brainstorming process led me to two options: either create another CLI interface for this, or start from scratch 
using what I learned to create a more robust version of the original project.

Clearly, I've opted for the second possibility. 

## Commands

```
Usage
  bloghead [command]

Available commands:
  add       Add a new bloghead element (template, datatype)
  create    Create a new page
  publish   Build the static site 

```
