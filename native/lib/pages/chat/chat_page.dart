import 'package:flutter/material.dart';
import 'package:flutter_hooks/flutter_hooks.dart';
import 'package:hooks_riverpod/hooks_riverpod.dart';
import 'package:myapp/pages/chat/chat_model.dart';
import 'package:myapp/pages/chat/message_body.dart';
import 'package:myapp/pages/chat/show_message_form.dart';
import 'package:myapp/widgets/header/header_widget.dart';

final messagesProvider = StateNotifierProvider((_) => MessagesState());

class ChatPage extends HookWidget {
  @override
  Widget build(context) {
    final messagesState = useProvider(messagesProvider.state);

    return Scaffold(
      appBar: Header('Chat'),
      body: messagesState.when(
        data: (messages) => RefreshIndicator(
          onRefresh: () => context.read(messagesProvider).getMessages(),
          child: ListView.builder(
            padding: const EdgeInsets.only(bottom: 80.0),
            itemCount: messages.length,
            itemBuilder: (context, position) => MessageBody(
              messages[position],
              context.read(messagesProvider).patchMessage,
              context.read(messagesProvider).deleteMessage,
            ),
          ),
        ),
        loading: () => Center(child: CircularProgressIndicator()),
        error: (_, __) => Center(
          child: Column(
            mainAxisSize: MainAxisSize.min,
            children: [
              Text('データの取得に失敗しました。'),
              SizedBox(height: 8),
              ElevatedButton.icon(
                icon: Icon(
                  Icons.refresh,
                  color: Colors.black,
                ),
                label: Text(
                  '再度読み込み',
                  style: TextStyle(color: Colors.black),
                ),
                style: ElevatedButton.styleFrom(
                  primary: Colors.white,
                  shape: RoundedRectangleBorder(
                    borderRadius: BorderRadius.circular(10),
                  ),
                ),
                onPressed: () => context.read(messagesProvider).getMessages(),
              ),
            ],
          ),
        ),
      ),
      floatingActionButton: FloatingActionButton(
        child: Icon(Icons.add),
        tooltip: 'メッセージを送る',
        onPressed: () => showNewMessageForm(context, context.read(messagesProvider).postMessage),
      ),
    );
  }
}
