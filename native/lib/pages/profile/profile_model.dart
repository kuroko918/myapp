import 'package:flutter_easyloading/flutter_easyloading.dart';
import 'package:hooks_riverpod/hooks_riverpod.dart';
import 'package:myapp/domains/user.dart';
import 'package:myapp/utils/api/user_api_client.dart';
import 'package:myapp/utils/firebase/firebase_auth.dart';

class UserState extends StateNotifier<AsyncValue<User>> {
  UserState() : super(const AsyncValue.loading()) {
    _createUserApiClient().then((_) => getUser(currentUser().uid));
  }

  UserApiClient api;
  User user;

  Future _createUserApiClient() async {
    api = await UserApiClient.create();
  }

  Future getUser(String userId) async {
    try {
      user = await api.getUser(userId);
      state = AsyncValue.data(user);
    } on Exception catch (e) {
      state = AsyncValue.error(e);
    }
  }

  Future putUser(String userId, User newUser) async {
    try {
      user = await api.putUser(userId, newUser);
      state = AsyncValue.data(user);
    } on Exception catch (e) {
      EasyLoading.showError('プロフィールの更新に失敗しました。');
      print(e);
    }
  }
}
