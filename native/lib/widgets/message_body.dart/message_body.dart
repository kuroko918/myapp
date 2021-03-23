import 'package:flutter/material.dart';
import 'package:flutter_hooks/flutter_hooks.dart';
import 'package:intl/intl.dart';
import 'package:myapp/domains/message.dart';
import 'package:myapp/utils/common/dialog.dart';
import 'package:myapp/utils/common/show_message_form.dart';
import 'package:myapp/utils/firebase/firebase_auth.dart';
import 'package:myapp/widgets/common/loading_box.dart';

class MessageBody extends HookWidget {
  MessageBody(this.message, this.patchMessage, this.deleteMessage);

  final Message message;
  final Future Function(String, Map<String, dynamic>) patchMessage;
  final Future Function(String) deleteMessage;

  @override
  Widget build(context) {
    renderMyMessageWidget() {
      return ListTile(
        leading: Container(
          width: 48,
          height: 48,
          child: Image.network(
            '${message.user.avatar}',
            fit: BoxFit.cover,
            loadingBuilder: (_, child, loadingProgress) {
              if (loadingProgress == null) return child;

              return showLoadingBox(48, 48);
            },
            errorBuilder: (_, __, ___) => Image.asset('images/flutter.png'),
          ),
        ),
        title: Text('${message.content}'),
        subtitle: message.createdAt == message.updatedAt ?
          Text("${DateFormat('yyyy-MM-dd HH:mm').format(message.createdAt.toLocal())}")
          : Text("${DateFormat('yyyy-MM-dd HH:mm').format(message.createdAt.toLocal())}（編集済）"),
      );
    }

    renderOtherMessageWidget() {
      return ListTile(
        trailing: Container(
          width: 48,
          height: 48,
          child: Image.network(
            '${message.user.avatar}',
            fit: BoxFit.cover,
            loadingBuilder: (_, child, loadingProgress) {
              if (loadingProgress == null) return child;

              return showLoadingBox(48, 48);
            },
            errorBuilder: (_, __, ___) => Image.asset('images/flutter.png'),
          ),
        ),
        title: Text('${message.content}'),
        subtitle: message.createdAt == message.updatedAt ?
          Text("${DateFormat('yyyy-MM-dd HH:mm').format(message.createdAt.toLocal())}")
          : Text("${DateFormat('yyyy-MM-dd HH:mm').format(message.createdAt.toLocal())}（編集済）"),
      );
    }

    renderRow() {
      if (currentUser().uid != message.userId) return null;

      return Row(
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
      );
    }

    renderColumnChildren() {
      List<Widget> children = [];
      if (currentUser().uid != message.userId) {
        children.add(renderOtherMessageWidget());
        return children;
      }

      children.add(renderMyMessageWidget());
      children.add(renderRow());

      return children;
    }

    return Card(
      elevation: 2.0,
      margin: EdgeInsets.all(8.0),
      child: Column(
        mainAxisSize: MainAxisSize.min,
        children: renderColumnChildren(),
      ),
    );
  }
}
