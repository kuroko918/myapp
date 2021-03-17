// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies

part of 'user.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;
User _$UserFromJson(Map<String, dynamic> json) {
  return _User.fromJson(json);
}

/// @nodoc
class _$UserTearOff {
  const _$UserTearOff();

// ignore: unused_element
  _User call(
      {String authToken,
      String id,
      String name,
      String email,
      String avatar,
      DateTime createdAt,
      DateTime updatedAt}) {
    return _User(
      authToken: authToken,
      id: id,
      name: name,
      email: email,
      avatar: avatar,
      createdAt: createdAt,
      updatedAt: updatedAt,
    );
  }

// ignore: unused_element
  User fromJson(Map<String, Object> json) {
    return User.fromJson(json);
  }
}

/// @nodoc
// ignore: unused_element
const $User = _$UserTearOff();

/// @nodoc
mixin _$User {
  String get authToken;
  String get id;
  String get name;
  String get email;
  String get avatar;
  DateTime get createdAt;
  DateTime get updatedAt;

  Map<String, dynamic> toJson();
  @JsonKey(ignore: true)
  $UserCopyWith<User> get copyWith;
}

/// @nodoc
abstract class $UserCopyWith<$Res> {
  factory $UserCopyWith(User value, $Res Function(User) then) =
      _$UserCopyWithImpl<$Res>;
  $Res call(
      {String authToken,
      String id,
      String name,
      String email,
      String avatar,
      DateTime createdAt,
      DateTime updatedAt});
}

/// @nodoc
class _$UserCopyWithImpl<$Res> implements $UserCopyWith<$Res> {
  _$UserCopyWithImpl(this._value, this._then);

  final User _value;
  // ignore: unused_field
  final $Res Function(User) _then;

  @override
  $Res call({
    Object authToken = freezed,
    Object id = freezed,
    Object name = freezed,
    Object email = freezed,
    Object avatar = freezed,
    Object createdAt = freezed,
    Object updatedAt = freezed,
  }) {
    return _then(_value.copyWith(
      authToken: authToken == freezed ? _value.authToken : authToken as String,
      id: id == freezed ? _value.id : id as String,
      name: name == freezed ? _value.name : name as String,
      email: email == freezed ? _value.email : email as String,
      avatar: avatar == freezed ? _value.avatar : avatar as String,
      createdAt:
          createdAt == freezed ? _value.createdAt : createdAt as DateTime,
      updatedAt:
          updatedAt == freezed ? _value.updatedAt : updatedAt as DateTime,
    ));
  }
}

/// @nodoc
abstract class _$UserCopyWith<$Res> implements $UserCopyWith<$Res> {
  factory _$UserCopyWith(_User value, $Res Function(_User) then) =
      __$UserCopyWithImpl<$Res>;
  @override
  $Res call(
      {String authToken,
      String id,
      String name,
      String email,
      String avatar,
      DateTime createdAt,
      DateTime updatedAt});
}

/// @nodoc
class __$UserCopyWithImpl<$Res> extends _$UserCopyWithImpl<$Res>
    implements _$UserCopyWith<$Res> {
  __$UserCopyWithImpl(_User _value, $Res Function(_User) _then)
      : super(_value, (v) => _then(v as _User));

  @override
  _User get _value => super._value as _User;

  @override
  $Res call({
    Object authToken = freezed,
    Object id = freezed,
    Object name = freezed,
    Object email = freezed,
    Object avatar = freezed,
    Object createdAt = freezed,
    Object updatedAt = freezed,
  }) {
    return _then(_User(
      authToken: authToken == freezed ? _value.authToken : authToken as String,
      id: id == freezed ? _value.id : id as String,
      name: name == freezed ? _value.name : name as String,
      email: email == freezed ? _value.email : email as String,
      avatar: avatar == freezed ? _value.avatar : avatar as String,
      createdAt:
          createdAt == freezed ? _value.createdAt : createdAt as DateTime,
      updatedAt:
          updatedAt == freezed ? _value.updatedAt : updatedAt as DateTime,
    ));
  }
}

@JsonSerializable()

/// @nodoc
class _$_User with DiagnosticableTreeMixin implements _User {
  const _$_User(
      {this.authToken,
      this.id,
      this.name,
      this.email,
      this.avatar,
      this.createdAt,
      this.updatedAt});

  factory _$_User.fromJson(Map<String, dynamic> json) =>
      _$_$_UserFromJson(json);

  @override
  final String authToken;
  @override
  final String id;
  @override
  final String name;
  @override
  final String email;
  @override
  final String avatar;
  @override
  final DateTime createdAt;
  @override
  final DateTime updatedAt;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'User(authToken: $authToken, id: $id, name: $name, email: $email, avatar: $avatar, createdAt: $createdAt, updatedAt: $updatedAt)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'User'))
      ..add(DiagnosticsProperty('authToken', authToken))
      ..add(DiagnosticsProperty('id', id))
      ..add(DiagnosticsProperty('name', name))
      ..add(DiagnosticsProperty('email', email))
      ..add(DiagnosticsProperty('avatar', avatar))
      ..add(DiagnosticsProperty('createdAt', createdAt))
      ..add(DiagnosticsProperty('updatedAt', updatedAt));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _User &&
            (identical(other.authToken, authToken) ||
                const DeepCollectionEquality()
                    .equals(other.authToken, authToken)) &&
            (identical(other.id, id) ||
                const DeepCollectionEquality().equals(other.id, id)) &&
            (identical(other.name, name) ||
                const DeepCollectionEquality().equals(other.name, name)) &&
            (identical(other.email, email) ||
                const DeepCollectionEquality().equals(other.email, email)) &&
            (identical(other.avatar, avatar) ||
                const DeepCollectionEquality().equals(other.avatar, avatar)) &&
            (identical(other.createdAt, createdAt) ||
                const DeepCollectionEquality()
                    .equals(other.createdAt, createdAt)) &&
            (identical(other.updatedAt, updatedAt) ||
                const DeepCollectionEquality()
                    .equals(other.updatedAt, updatedAt)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^
      const DeepCollectionEquality().hash(authToken) ^
      const DeepCollectionEquality().hash(id) ^
      const DeepCollectionEquality().hash(name) ^
      const DeepCollectionEquality().hash(email) ^
      const DeepCollectionEquality().hash(avatar) ^
      const DeepCollectionEquality().hash(createdAt) ^
      const DeepCollectionEquality().hash(updatedAt);

  @JsonKey(ignore: true)
  @override
  _$UserCopyWith<_User> get copyWith =>
      __$UserCopyWithImpl<_User>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$_$_UserToJson(this);
  }
}

abstract class _User implements User {
  const factory _User(
      {String authToken,
      String id,
      String name,
      String email,
      String avatar,
      DateTime createdAt,
      DateTime updatedAt}) = _$_User;

  factory _User.fromJson(Map<String, dynamic> json) = _$_User.fromJson;

  @override
  String get authToken;
  @override
  String get id;
  @override
  String get name;
  @override
  String get email;
  @override
  String get avatar;
  @override
  DateTime get createdAt;
  @override
  DateTime get updatedAt;
  @override
  @JsonKey(ignore: true)
  _$UserCopyWith<_User> get copyWith;
}
