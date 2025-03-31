package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"time"
)

type Arguments struct {
	directory  string
	title      string
	categories string
	template   string
}

func parseArgument(arg string, args *Arguments) {
	if strings.Contains(arg, "--dir=") {
		args.directory = strings.ReplaceAll(arg, "--dir=", "")

		if args.directory[len(args.directory)-1] != '/' {
			args.directory = args.directory + "/"
		}

	}
	if strings.Contains(arg, "--title=") {
		args.title = strings.ReplaceAll(arg, "--title=", "")
	}
	if strings.Contains(arg, "--categories=") {
		args.categories = strings.ReplaceAll(arg, "--categories=", "")
	}
	if strings.Contains(arg, "--template=") {
		args.template = strings.ReplaceAll(arg, "--template=", "")
	}
  if strings.Contains(arg, "--help") {
    fmt.Println("Usage: entry --dir=<directory> --title=<title> --categories=<categories> --template=<template>")
    os.Exit(0)
  }
}

func getArguments() Arguments {
	home, err := os.UserHomeDir()

	if err != nil {
		panic(err)
	}

	args := Arguments{
		directory:  "_posts/",
		title:      "<template>",
		categories: "",
		template:   home + "/.config/entry/template.markdown",
	}

	for i := 1; i < len(os.Args); i++ {
		parseArgument(os.Args[i], &args)
	}

	return args
}

func main() {
	args := getArguments()
	template, err := os.ReadFile(args.template)

	if err != nil {
		fmt.Fprintf(os.Stderr, "An error happend while reading the template")
		panic(err)
	}

	now := time.Now()
	date := now.Format("2006-01-02 15:04:05 -0700")

	templateTitle := fmt.Sprintf("%s%s-%s.markdown",
		args.directory,
		now.Format("2006-01-02"),
		args.title,
	)

	template = bytes.ReplaceAll(template, []byte("{title}"), []byte(args.title))
	template = bytes.ReplaceAll(template, []byte("{date}"), []byte(date))
	template = bytes.ReplaceAll(template, []byte("{categories}"), []byte(args.categories))
	fmt.Println(string(template))

	err = os.WriteFile(templateTitle, template, 0644)

	if err != nil {
		panic(err)
	}
}
