import 'package:flutter/material.dart';
import 'package:flutter_hooks/flutter_hooks.dart';
import 'package:hooks_riverpod/hooks_riverpod.dart';
import 'package:myapp/domains/user.dart';
import 'package:myapp/pages/profile/profile_model.dart';
import 'package:myapp/utils/firebase/firebase_auth.dart';
import 'package:myapp/widgets/common/loading_box.dart';
import 'package:myapp/widgets/header/header_widget.dart';
import 'package:myapp/widgets/profile_edit_form/profile_edit_form.dart';

final userProvider = StateNotifierProvider((_) => UserState());

class ProfilePage extends HookWidget {
  @override
  Widget build(context) {
    final userState = useProvider(userProvider.state);

    _renderDataWidget(User user) {
      return Center(
        child: SingleChildScrollView(
          child: Column(
            mainAxisSize: MainAxisSize.min,
            children: [
              ClipRRect(
                borderRadius: BorderRadius.circular(100),
                child: Image.network(
                  '${user.avatar}',
                  width: 150,
                  height: 150,
                  fit: BoxFit.cover,
                  loadingBuilder: (_, child, loadingProgress) {
                    if (loadingProgress == null) return child;

                    return showLoadingBox(150, 150);
                  },
                  errorBuilder: (_, __, ___) {
                    return Image.asset(
                      'images/flutter.png',
                      width: 150,
                      height: 150,
                      fit: BoxFit.cover,
                    );
                  },
                ),
              ),
              SizedBox(height: 8),
              Text(
                '${user.name}',
                style: TextStyle(
                  fontSize: 20,
                  fontWeight: FontWeight.bold,
                ),
              ),
              SizedBox(height: 8),
              ElevatedButton.icon(
                icon: Icon(Icons.edit),
                label: Text('プロフィールを編集'),
                style: ElevatedButton.styleFrom(
                  shape: RoundedRectangleBorder(
                    borderRadius: BorderRadius.circular(10),
                  ),
                ),
                onPressed: () {
                  Navigator.push(
                    context,
                    MaterialPageRoute(builder: (context) => ProfileEditForm(user, context.read(userProvider).putUser)),
                  );
                }
              ),
            ],
          ),
        ),
      );
    }

    _renderErrorWidget() {
      return Center(
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
      );
    }

    return Scaffold(
      appBar: Header('Profile'),
      body: userState.when(
        data: (user) => _renderDataWidget(user),
        loading: () => Center(child: CircularProgressIndicator()),
        error: (_, __) => _renderErrorWidget(),
      ),
    );
  }
}
