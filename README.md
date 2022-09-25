# HLive In Wails Prototype

## About

An experiment in getting HLive to work in Wails

## Progress

Adapted from the Wails vanilla JS template. 

- Both wails dev and wails build work.
- Attempted to make the prototype match the 
- File upload not yet supported, it's going to be rewritten in HLive soon
- HLive doesn't integrate with Vite. 
  - We needed to make is so the development frontend asset files mirror the production files.
  - See vite.config.js
- HLive will detect if it's running in Wails. 
  - This was added to the 0.2.0 HLive branch, an upcoming version.
- Events are used to communicate between the browser and HLive.
- No need to bind any functions 
- Logging is set to debug levels 
- Can't get the page logger for the PageServer
- The index.html file is used for Vite asset building
- The index.html will redirect to hlive.html
- An asset handler is used to render hlive.html
- Wails ipc and runtime are auto-injected into the page by the page server
- Page asset management is rough and still done in the app for now
  - An area for improvement
- The interplay of Wails and HLive context is a little awkward
  - Docs will be needed to explain when each is relevant
  - Most of the time HLive is using the Wails context which is ideal
- When in wails dev, the app and the browser share the same page 

## Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`.
