// Copyright 2020 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bertqa

import "html/template"

const style template.CSS = `
html, body { height: 100vh; }

body {
	margin: 0;
	font-family: sans-serif;
	color: #333333;
	overflow: auto;
}

a {
	color: inherit;
	text-decoration: none;
	outline: none;
}

textarea,
input,
button {
	border: 0;
	outline: none;
	font-family: inherit;
	font-size: inherit;
	color: inherit;
	text-align: left;
}

#answers {
	width: 25rem;
}

#answers button.active,
#highlightable-text span.active {
	background-color: #8fe5fa;
}

#loader {
	align-self: center;
	border: 3px solid transparent;
	border-top: 3px solid #0aadd8;
	border-radius: 50%;
	width: 1rem;
	height: 1rem;
	animation: spin 2s linear infinite;
}

@keyframes spin {
	0% { transform: rotate(0deg); }
	100% { transform: rotate(360deg); }
}

/* Style classes loosely based on Tailwind */

.flex { display: flex; }
.flex-col { flex-direction: column; }
.flex-grow { flex-grow: 1; }

.p-2 { padding: .5rem; }
.py-2 { padding-top: .5rem; padding-bottom: .5rem; }
.px-4 { padding-left: 1rem; padding-right: 1rem; }
.mt-2 { margin-top: .5rem; }
.mb-2 { margin-bottom: .5rem; }
.mr-2 { margin-right: .5rem; }

.text-xl { font-size: 1.25rem; }
.font-bold { font-weight: bold; }
.italic { font-style: italic; }
.text-blue { color: #0aadd8; }
.text-red { color: #f56565; }
.text-gray { color: #7e8997; }
.text-transparent { color: transparent; }

.bg-white { background-color: #fff; }
.bg-gray-200 { background-color: #edf2f7; }
.bg-gray-300 { background-color: #e2e8f0; }
.bg-transparent { background-color: transparent; }
.bg-blue { background-color: #8fe5fa; }
.hover\:bg-light-blue:hover { background-color: #9ff5ff; }
.hover\:bg-gray-200:hover { background-color: #edf2f7; }

.shadow { box-shadow: 0 0 5px 0 rgba(0,0,0,0.1); }
.z-10 { z-index: 10; }
.rounded { border-radius: 0.25rem; }
.resize-none { resize: none; }
.absolute { position: absolute; }
.relative { position: relative; }
.block { display: block; }
.hidden { display: none; }
.cursor-pointer { cursor: pointer; }
.inset-0 { top: 0; right: 0; bottom: 0; left: 0; }
.overflow-hidden { overflow: hidden; }
.overflow-auto { overflow: auto; }
`
