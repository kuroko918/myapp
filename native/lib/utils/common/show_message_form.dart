import 'package:flutter/material.dart';
import 'package:myapp/domains/message.dart';

Future<void> showNewMessageForm(BuildContext context, Future Function(Map<String, dynamic>) postMessage) {
  return showDialog(
    context: context,
    barrierDismissible: false,
    builder: (context) {
      String editingContent = '';

      return AlertDialog(
        title: Text('新しいメッセージを入力'),
        content: Row(
          children: [
            Expanded(
              child: TextField(
                autofocus: true,
                decoration: InputDecoration(hintText: 'メッセージ'),
                onChanged: (value) {
                  editingContent = value;
                },
              ),
            ),
          ],
        ),
        actions: [
          TextButton(
            child: Text('送信'),
            onPressed: () async {
              final submitContent = editingContent.trim();

              if (submitContent == '') return;
              if (submitContent.length > 140) return;

              await postMessage({'content': editingContent});
              Navigator.of(context).pop();
            }
          ),
          TextButton(
            child: Text('キャンセル'),
            onPressed: () => Navigator.of(context).pop(),
          ),
        ],
      );
    }
  );
}

Future<void> showEditMessageForm(BuildContext context, Message message, Future Function(String, Map<String, dynamic>) patchMessage) {
  return showDialog(
    context: context,
    barrierDismissible: false,
    builder: (context) {
      String editingContent = message.content;

      return AlertDialog(
        title: Text('メッセージを編集'),
        content: Row(
          children: [
            Expanded(
              child: TextField(
                autofocus: true,
                controller: TextEditingController(text: message.content),
                decoration: InputDecoration(hintText: 'メッセージ'),
                onChanged: (value) {
                  editingContent = value;
                },
              ),
            ),
          ],
        ),
        actions: [
          TextButton(
            child: Text('送信'),
            onPressed: () async {
              final submitContent = editingContent.trim();

              if (submitContent == '') return;
              if (submitContent == message.content) return;
              if (submitContent.length > 140) return;

              await patchMessage(message.id, {'content': submitContent});
              Navigator.of(context).pop();
            }
          ),
          TextButton(
            child: Text('キャンセル'),
            onPressed: () => Navigator.of(context).pop(),
          ),
        ],
      );
    }
  );
}
