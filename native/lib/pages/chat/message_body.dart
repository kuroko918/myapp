import 'package:flutter/material.dart';
import 'package:flutter_hooks/flutter_hooks.dart';
import 'package:intl/intl.dart';
import 'package:myapp/domains/message.dart';
import 'package:myapp/pages/chat/show_message_form.dart';
import 'package:myapp/utils/common/dialog.dart';

class MessageBody extends HookWidget {
  MessageBody(this.message, this.patchMessage, this.deleteMessage);

  final Message message;
  final Future Function(String, Map<String, dynamic>) patchMessage;
  final Future Function(String) deleteMessage;

  @override
  Widget build(context) {
    return Card(
      elevation: 2.0,
      margin: EdgeInsets.all(8.0),
      child: Column(
        mainAxisSize: MainAxisSize.min,
        children: [
          ListTile(
            leading: Container(
              width: 48,
              height: 48,
              child: Image.network(
                '${message.user.avatar}',
                errorBuilder: (_, __, ___) => Image.asset('images/flutter.png'),
              ),
            ),
            title: Text('${message.content}'),
            subtitle: message.createdAt == message.updatedAt ?
              Text("${DateFormat('yyyy-MM-dd HH:mm').format(message.createdAt.toLocal())}")
              : Text("${DateFormat('yyyy-MM-dd HH:mm').format(message.createdAt.toLocal())}（編集済）"),
          ),
          Row(
            mainAxisAlignment: MainAxisAlignment.end,
            children: [
              TextButton(
                child: Text('編集'),
                onPressed: () => showEditMessageForm(context, message, patchMessage),
              ),
              SizedBox(width: 8),
              TextButton(
                child: Text('削除'),
                onPressed: () async {
                  final result = await showAlertDialog(context, 'このメッセージを削除しますか？', ['削除', 'キャンセル']);
                  if (result) deleteMessage(message.id);
                }
              ),
              SizedBox(width: 8),
            ],
          ),
        ],
      ),
    );
  }
}
