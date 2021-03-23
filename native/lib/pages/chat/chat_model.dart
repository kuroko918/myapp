import 'package:flutter_easyloading/flutter_easyloading.dart';
import 'package:hooks_riverpod/hooks_riverpod.dart';
import 'package:myapp/domains/message.dart';
import 'package:myapp/utils/api/message_api_client.dart';

class MessagesState extends StateNotifier<AsyncValue<List<Message>>> {
  MessagesState() : super(const AsyncValue.loading()) {
    _createMessageApiClient().then((_) => getMessages());
  }

  MessageApiClient api;
  List<Message> messageList = [];

  Future _createMessageApiClient() async {
    api = await MessageApiClient.create();
  }

  Future getMessages() async {
    try {
      final messages = await api.getMessages();
      messageList = messages['messages'];
      state = AsyncValue.data(messageList);
    } on Exception catch (e) {
      state = AsyncValue.error(e);
    }
  }

  Future postMessage(Map<String, dynamic> params) async {
    try {
      final newMessage = await api.postMessage(params);
      messageList.add(newMessage);
      state = AsyncValue.data(messageList);
    } on Exception catch (e) {
      EasyLoading.showError('メッセージの作成に失敗しました。');
      print(e);
    }
  }

  Future patchMessage(String messageId, Map<String, dynamic> params) async {
    try {
      final newMessage = await api.patchMessage(messageId, params);
      final newMessageList = messageList.map((message) {
        if (message.id == newMessage.id) return newMessage;
        return message;
      }).toList();
      state = AsyncValue.data(newMessageList);
    } on Exception catch (e) {
      EasyLoading.showError('メッセージの更新に失敗しました。');
      print(e);
    }
  }

  Future deleteMessage(String messageId) async {
    try {
      await api.deleteMessage(messageId);
      messageList.removeWhere((Message m) => m.id == messageId);
      state = AsyncValue.data(messageList);
    } on Exception catch (e) {
      EasyLoading.showError('メッセージの削除に失敗しました。');
      print(e);
    }
  }
}
