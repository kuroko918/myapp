import 'dart:io';
import 'package:flutter/material.dart';
import 'package:flutter_easyloading/flutter_easyloading.dart';
import 'package:flutter_hooks/flutter_hooks.dart';
import 'package:myapp/domains/user.dart';
import 'package:myapp/utils/common/upload_image.dart';
import 'package:myapp/utils/common/show_get_image_modal.dart';

class ProfileEditForm extends HookWidget {
  ProfileEditForm(this.user, this.putMessage);

  final User user;
  final Future Function(String, User) putMessage;

  @override
  Widget build(context) {
    final _formKey = useMemoized(() => GlobalKey<FormState>());
    final cacheImage = useState(File(''));
    final avatar = useState(user.avatar);
    final name = useState(user.name);
    final email = useState(user.email);

    _handleSubmit() async {
      if (EasyLoading.isShow) return;
      if (!_formKey.currentState.validate()) return;

      EasyLoading.show(status: 'loading...');

      String imageURL;

      if (cacheImage.value.path != '') imageURL = await uploadImage(cacheImage.value, user.id);

      await putMessage(user.id, user.copyWith(
        avatar: imageURL ?? avatar.value,
        name: name.value,
        email: email.value,
      ));
      EasyLoading.dismiss();
      Navigator.of(context).pop();
    }

    renderAvatarImageField() {
      _onTap() async {
        final image = await showGetImageModal(context);
        if (image == null) return;

        cacheImage.value = image;
      }

      return Stack(
        alignment: AlignmentDirectional.bottomCenter,
        children: [
          GestureDetector(
            child: Center(
              child: ClipRRect(
                borderRadius: BorderRadius.circular(100),
                child: cacheImage.value.path != '' ?
                  Image.file(
                    cacheImage.value,
                    width: 120,
                    height: 120,
                    fit: BoxFit.cover,
                  )
                  : Image.network(
                    '${avatar.value}',
                    width: 120,
                    height: 120,
                    fit: BoxFit.cover,
                    errorBuilder: (_, __, ___) {
                      return Image.asset(
                        'images/flutter.png',
                        width: 120,
                        height: 120,
                        fit: BoxFit.cover,
                      );
                    },
                  ),
              ),
            ),
            onTap: () => _onTap(),
          ),
          Align(
            alignment: Alignment(0.1, 0),
            child: GestureDetector(
              child: ClipRRect(
                borderRadius: BorderRadius.circular(100),
                child: Icon(
                  Icons.add_circle_rounded,
                  size: 40,
                  color: Colors.blue,
                ),
              ),
              onTap: () => _onTap(),
            ),
          ),
        ],
      );
    }

    renderNameTextField() {
      return Padding(
        padding: EdgeInsets.symmetric(horizontal: 16),
        child: TextFormField(
          maxLength: 30,
          initialValue: '${user.name}',
          autovalidateMode: AutovalidateMode.onUserInteraction,
          decoration: InputDecoration(labelText: '名前'),
          validator: (value) {
            final submitName = value.trim();
            if (submitName.isEmpty) return '名前を入力してください';
            if (submitName.length > 30) return '名前は30文字以内で入力してください';

            return null;
          },
          onChanged: (value) => name.value = value,
        ),
      );
    }

    renderEmailTextField() {
      return Padding(
        padding: EdgeInsets.symmetric(horizontal: 16),
        child: TextFormField(
          initialValue: '${user.email}',
          decoration: InputDecoration(labelText: 'Email'),
          autovalidateMode: AutovalidateMode.disabled,
          validator: (value) {
            final submitEmail = value.trim();
            if (submitEmail.isEmpty) return 'Emailを入力してください';
            final emailRegx = r"[a-zA-Z0-9.a-zA-Z0-9.!#$%&'*+-/=?^_`{|}~]+@[a-zA-Z0-9]+\.[a-zA-Z]+";
            if (!RegExp(emailRegx).hasMatch(submitEmail)) return '適切なEmailを入力してください';

            return null;
          },
          onChanged: (value) => email.value = value,
        ),
      );
    }

    renderSubmitButton() {
      return Container(
        width: double.infinity,
        margin: EdgeInsets.symmetric(horizontal: 16),
        child: SizedBox(
          height: 50,
          child: ElevatedButton(
            child: Text('保存する'),
            style: ElevatedButton.styleFrom(
              shape: RoundedRectangleBorder(
                borderRadius: BorderRadius.circular(50),
              ),
            ),
            onPressed: () => _handleSubmit(),
          ),
        ),
      );
    }

    return Scaffold(
      appBar: AppBar(
        leading: IconButton(
          icon: Icon(Icons.close),
          onPressed: () => Navigator.of(context).pop(),
        ),
        title: Text('Edit Profile'),
      ),
      body: Form(
        key: _formKey,
        child: Center(
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.center,
            mainAxisSize: MainAxisSize.max,
            mainAxisAlignment: MainAxisAlignment.end,
            children: [
              SizedBox(height: 32),
              renderAvatarImageField(),
              renderNameTextField(),
              renderEmailTextField(),
              Spacer(),
              renderSubmitButton(),
              SizedBox(height: 24),
            ],
          ),
        ),
      ),
    );
  }
}
