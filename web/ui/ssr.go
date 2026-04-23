package ui

// CSS is the stylesheet for the tinywasm.app website.
const CSS = `
*, *::before, *::after { box-sizing: border-box; margin: 0; padding: 0; }

body {
	background: #0d1117;
	color: #e6edf3;
	font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, sans-serif;
	min-height: 100vh;
	display: flex;
	align-items: center;
	justify-content: center;
}

.uc-page {
	display: flex;
	align-items: center;
	justify-content: center;
	width: 100%;
	min-height: 100vh;
}

.uc-card {
	text-align: center;
	padding: 3rem 2rem;
	border: 1px solid #30363d;
	border-radius: 12px;
	background: #161b22;
	max-width: 480px;
	width: 90%;
}

.uc-title {
	font-size: 2rem;
	font-weight: 700;
	color: #58a6ff;
	letter-spacing: -0.02em;
	margin-bottom: 0.75rem;
}

.uc-sub {
	font-size: 1.1rem;
	font-weight: 600;
	color: #8b949e;
	margin-bottom: 1rem;
}

.uc-msg {
	font-size: 0.95rem;
	color: #6e7681;
	line-height: 1.6;
}
`
