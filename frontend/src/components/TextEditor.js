import React, { useEffect, useRef, useState } from 'react';
import { useEditor, EditorContent } from '@tiptap/react';
import StarterKit from '@tiptap/starter-kit';
import "./TextEditor.css";
import Bold from '@tiptap/extension-bold';
import Italic from '@tiptap/extension-italic';
import Strike from '@tiptap/extension-strike';
import Underline from '@tiptap/extension-underline';
import Code from '@tiptap/extension-code';
import Paragraph from '@tiptap/extension-paragraph';
import BulletList from '@tiptap/extension-bullet-list';
import OrderedList from '@tiptap/extension-ordered-list';
import ListItem from '@tiptap/extension-list-item';
import Link from '@tiptap/extension-link';
import CodeBlock from '@tiptap/extension-code-block';

const CustomParagraph = Paragraph.extend({
  addAttributes() {
    return {
      ...this.parent?.(),
      id: {
        default: null,
      },
    };
  },
});

const TextEditor = ({ title, content, setContent }) => {
  const editor = useEditor({
    extensions: [
      StarterKit.configure({
        history: false,
      }),
      Bold,
      Italic,
      Strike,
      Underline,
      Code,
      CustomParagraph,
      BulletList,
      OrderedList,
      ListItem,
      Link,
      CodeBlock,
    ],
    content: content,
    onUpdate: ({ editor }) => {
      console.log('update received');
      const text = editor.getHTML();
      setContent(text);
    },
  });

  const [prevContent, setPrevContent] = useState(content);
  const [currentContent, setCurrentContent] = useState('');

  const DEBOUNCE_DELAY1 = 500;
  let debounceTimer = null;

  const setEditorContent = () => {
    if (editor) {
      setCurrentContent(editor.getHTML());
    }
  };

  useEffect(() => {
    console.log('use effect triggered 2');
    if (editor && content) {
      editor.commands.setContent(content);
      setPrevContent(content);
    }
  }, [editor, content]);

  useEffect(() => {
    const handleKeyDown = () => {
      setEditorContent();
    };

    document.addEventListener('keydown', handleKeyDown);

    return () => {
      document.removeEventListener('keydown', handleKeyDown);
    };
  }, [editor]);

  function onButtonPress() {
  const currentText = currentContent;
  const prevText = prevContent;

  let i = 0;
  while (i < currentText.length && i < prevText.length && currentText[i] === prevText[i]) {
    i++;
  }

  const changes = currentText.slice(i);

  console.log('current content:', currentContent);
  console.log('total content:', prevContent);
  console.log('changes:', changes);

  if (debounceTimer) {
    clearTimeout(debounceTimer);
  }

  debounceTimer = setTimeout(() => {
    sendUpdate(changes);
    setPrevContent(currentContent);
  }, DEBOUNCE_DELAY1);
}

  const websocketRef = useRef(null);

  useEffect(() => {
    // Establish WebSocket connection with the backend
    console.log('use effect triggered 1');
    const ws = new WebSocket(`ws://localhost:8080/ws/${encodeURIComponent(title)}`);
    websocketRef.current = ws;

    ws.onopen = () => {
      console.log('WebSocket connection established');
    };

    ws.onmessage = (event) => {
      console.log("message received", event.data);
      const operation = JSON.parse(event.data);
      if (operation.type === 0) {
        const element = operation.element;
        const attrs = { id: element.id, position: element.position };
        const newContent = `<p ${Object.entries(attrs).map(([key, value]) => `${key}="${JSON.stringify(value)}"`).join(' ')}>${element.content}</p>`;
        console.log('new content weoowpeqodpwo[');
        console.log(newContent);
        editor.commands.insertContent(newContent);
      } else if (operation.type === 1) {
        const elementId = operation.element.id;
        editor.commands.command(({ tr }) => {
          const element = tr.doc.nodesBetween(0, tr.doc.content.size, (node, pos) => {
            if (node.attrs.id === elementId) {
              return true;
            }
          });
          if (element) {
            tr.delete(element.pos, element.pos + element.node.nodeSize);
          }
          return true;
        });
      }
    };

    ws.onerror = (error) => {
      console.error('WebSocket error:', error);
    };

    ws.onclose = () => {
      console.log('WebSocket connection closed');
    };

    return () => {
      ws.close();
    };
  }, [editor, title]);

  const DEBOUNCE_DELAY = 500;
  let updateTimeout = null;

  const sendUpdate = (newContent) => {
    if (websocketRef.current && websocketRef.current.readyState === WebSocket.OPEN) {
      const operation = {
        type: 0,
        element: {
          content: newContent,
        },
      };

      if (updateTimeout) {
        clearTimeout(updateTimeout);
      }

      updateTimeout = setTimeout(() => {
        websocketRef.current.send(JSON.stringify(operation));
      }, DEBOUNCE_DELAY);
    }
  };

  if (!editor) {
    return null;
  }

  return (
    <div>
      <div className="toolbar">
        <button onClick={() => editor.chain().focus().toggleBold().run()} className={editor.isActive('bold') ? 'is-active' : ''}>
          Bold
        </button>
        <button onClick={() => editor.chain().focus().toggleItalic().run()} className={editor.isActive('italic') ? 'is-active' : ''}>
          Italic
        </button>
        <button onClick={() => editor.chain().focus().toggleStrike().run()} className={editor.isActive('strike') ? 'is-active' : ''}>
          Strike
        </button>
        <button onClick={() => editor.chain().focus().toggleUnderline().run()} className={editor.isActive('underline') ? 'is-active' : ''}>
          Underline
        </button>
        <button onClick={() => editor.chain().focus().toggleCode().run()} className={editor.isActive('code') ? 'is-active' : ''}>
          Code
        </button>
        <button onClick={() => editor.chain().focus().setParagraph().run()} className={editor.isActive('paragraph') ? 'is-active' : ''}>
          Paragraph
        </button>
        <button onClick={() => editor.chain().focus().toggleBulletList().run()} className={editor.isActive('bulletList') ? 'is-active' : ''}>
          Bullet List
        </button>
        <button onClick={() => editor.chain().focus().toggleOrderedList().run()} className={editor.isActive('orderedList') ? 'is-active' : ''}>
          Ordered List
        </button>
        <button onClick={() => editor.chain().focus().toggleCodeBlock().run()} className={editor.isActive('codeBlock') ? 'is-active' : ''}>
          Code Block
        </button>
      </div>
      <h1>{title}</h1>
      <button onClick={onButtonPress}>
        HELOOOOOOOOOOOOO
      </button>
      <EditorContent editor={editor} className="editor-content" />
    </div>
  );
};

export default TextEditor;

//https://tiptap.dev/docs/editor/examples/default