package tui

var LateTheme = []byte(`
{
  "document": {
    "block_prefix": "",
    "block_suffix": "",
    "color": "#F3F4F6",
    "background_color": "#0E0E10",
    "margin": 0
  },
  "paragraph": {
    "margin": 0,
    "background_color": "#0E0E10"
  },
  "block_quote": {
    "indent": 1,
    "indent_token": "│ ",
    "color": "#8A94A6",
    "background_color": "#0E0E10"
  },
  "list": {
    "level_indent": 2,
    "background_color": "#0E0E10"
  },
  "bullet": {
    "color": "#E5A85C"
  },
  "enumeration": {
    "color": "#E5A85C",
    "block_suffix": ". "
  },
  "task": {
    "ticked": "[x] ",
    "unticked": "[ ] ",
    "color": "#E5A85C"
  },
  "heading": {
    "block_suffix": "\n",
    "color": "#E5A85C",
    "bold": true
  },
  "h1": {
    "prefix": "# "
  },
  "h2": {
    "prefix": "## "
  },
  "h3": {
    "prefix": "### "
  },
  "strong": {
    "bold": true,
    "color": "#E5A85C"
  },
  "emph": {
    "italic": true,
    "color": "#62B3D5"
  },
  "code": {
    "prefix": " ",
    "suffix": " ",
    "color": "#62B3D5",
    "background_color": "#1B1B1E"
  },
  "code_block": {
    "margin": 0,
    "chroma": {
      "background": {
        "background_color": "#141416"
      },
      "text": {
        "color": "#F3F4F6",
        "background_color": "#141416"
      },
      "error": {
        "color": "#EF4444",
        "background_color": "#141416"
      },
      "comment": {
        "color": "#64748B"
      },
      "keyword": {
        "color": "#E5A85C"
      },
      "literal": {
        "color": "#62B3D5"
      },
      "name_tag": {
        "color": "#62B3D5"
      },
      "operator": {
        "color": "#F3F4F6"
      },
      "string": {
        "color": "#A7C080"
      }
    },
    "background_color": "#141416"
  },
  "table": {
    "center": false,
    "margin": 0,
    "color": "#F3F4F6",
    "background_color": "#0E0E10"
  },
  "table_header": {
    "color": "#E5A85C",
    "background_color": "#0E0E10",
    "bold": true
  },
  "table_cell": {
    "color": "#F3F4F6",
    "background_color": "#0E0E10"
  },
  "link": {
    "color": "#62B3D5",
    "underline": true
  },
  "image": {
    "color": "#62B3D5",
    "underline": true
  }
}
`)
