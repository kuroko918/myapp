import 'package:flutter/material.dart';
import 'package:flutter_hooks/flutter_hooks.dart';
import 'package:hooks_riverpod/hooks_riverpod.dart';
import 'package:myapp/pages/profile/profile_model.dart';
import 'package:myapp/utils/firebase/firebase_auth.dart';
import 'package:myapp/widgets/header/header_widget.dart';

final userProvider = StateNotifierProvider((_) => UserState());

class ProfilePage extends HookWidget {
  @override
  Widget build(context) {
    final userState = useProvider(userProvider.state);

    return Scaffold(
      appBar: Header('Profile'),
      body: userState.when(
        data: (user) => Center(
          child: Column(
            mainAxisSize: MainAxisSize.min,
            children: [
              Image.network('${user.avatar}'),
              Text('${user.name}'),
            ],
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
                onPressed: () => context.read(userProvider).getUser(currentUser().uid),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
