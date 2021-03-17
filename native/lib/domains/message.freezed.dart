// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies

part of 'message.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;
Message _$MessageFromJson(Map<String, dynamic> json) {
  return _Message.fromJson(json);
}

/// @nodoc
class _$MessageTearOff {
  const _$MessageTearOff();

// ignore: unused_element
  _Message call(
      {String id,
      String content,
      String userId,
      DateTime createdAt,
      DateTime updatedAt,
      @JsonKey(name: 'user') User user}) {
    return _Message(
      id: id,
      content: content,
      userId: userId,
      createdAt: createdAt,
      updatedAt: updatedAt,
      user: user,
    );
  }

// ignore: unused_element
  Message fromJson(Map<String, Object> json) {
    return Message.fromJson(json);
  }
}

/// @nodoc
// ignore: unused_element
const $Message = _$MessageTearOff();

/// @nodoc
mixin _$Message {
  String get id;
  String get content;
  String get userId;
  DateTime get createdAt;
  DateTime get updatedAt;
  @JsonKey(name: 'user')
  User get user;

  Map<String, dynamic> toJson();
  @JsonKey(ignore: true)
  $MessageCopyWith<Message> get copyWith;
}

/// @nodoc
abstract class $MessageCopyWith<$Res> {
  factory $MessageCopyWith(Message value, $Res Function(Message) then) =
      _$MessageCopyWithImpl<$Res>;
  $Res call(
      {String id,
      String content,
      String userId,
      DateTime createdAt,
      DateTime updatedAt,
      @JsonKey(name: 'user') User user});

  $UserCopyWith<$Res> get user;
}

/// @nodoc
class _$MessageCopyWithImpl<$Res> implements $MessageCopyWith<$Res> {
  _$MessageCopyWithImpl(this._value, this._then);

  final Message _value;
  // ignore: unused_field
  final $Res Function(Message) _then;

  @override
  $Res call({
    Object id = freezed,
    Object content = freezed,
    Object userId = freezed,
    Object createdAt = freezed,
    Object updatedAt = freezed,
    Object user = freezed,
  }) {
    return _then(_value.copyWith(
      id: id == freezed ? _value.id : id as String,
      content: content == freezed ? _value.content : content as String,
      userId: userId == freezed ? _value.userId : userId as String,
      createdAt:
          createdAt == freezed ? _value.createdAt : createdAt as DateTime,
      updatedAt:
          updatedAt == freezed ? _value.updatedAt : updatedAt as DateTime,
      user: user == freezed ? _value.user : user as User,
    ));
  }

  @override
  $UserCopyWith<$Res> get user {
    if (_value.user == null) {
      return null;
    }
    return $UserCopyWith<$Res>(_value.user, (value) {
      return _then(_value.copyWith(user: value));
    });
  }
}

/// @nodoc
abstract class _$MessageCopyWith<$Res> implements $MessageCopyWith<$Res> {
  factory _$MessageCopyWith(_Message value, $Res Function(_Message) then) =
      __$MessageCopyWithImpl<$Res>;
  @override
  $Res call(
      {String id,
      String content,
      String userId,
      DateTime createdAt,
      DateTime updatedAt,
      @JsonKey(name: 'user') User user});

  @override
  $UserCopyWith<$Res> get user;
}

/// @nodoc
class __$MessageCopyWithImpl<$Res> extends _$MessageCopyWithImpl<$Res>
    implements _$MessageCopyWith<$Res> {
  __$MessageCopyWithImpl(_Message _value, $Res Function(_Message) _then)
      : super(_value, (v) => _then(v as _Message));

  @override
  _Message get _value => super._value as _Message;

  @override
  $Res call({
    Object id = freezed,
    Object content = freezed,
    Object userId = freezed,
    Object createdAt = freezed,
    Object updatedAt = freezed,
    Object user = freezed,
  }) {
    return _then(_Message(
      id: id == freezed ? _value.id : id as String,
      content: content == freezed ? _value.content : content as String,
      userId: userId == freezed ? _value.userId : userId as String,
      createdAt:
          createdAt == freezed ? _value.createdAt : createdAt as DateTime,
      updatedAt:
          updatedAt == freezed ? _value.updatedAt : updatedAt as DateTime,
      user: user == freezed ? _value.user : user as User,
    ));
  }
}

@JsonSerializable()

/// @nodoc
class _$_Message with DiagnosticableTreeMixin implements _Message {
  const _$_Message(
      {this.id,
      this.content,
      this.userId,
      this.createdAt,
      this.updatedAt,
      @JsonKey(name: 'user') this.user});

  factory _$_Message.fromJson(Map<String, dynamic> json) =>
      _$_$_MessageFromJson(json);

  @override
  final String id;
  @override
  final String content;
  @override
  final String userId;
  @override
  final DateTime createdAt;
  @override
  final DateTime updatedAt;
  @override
  @JsonKey(name: 'user')
  final User user;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'Message(id: $id, content: $content, userId: $userId, createdAt: $createdAt, updatedAt: $updatedAt, user: $user)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'Message'))
      ..add(DiagnosticsProperty('id', id))
      ..add(DiagnosticsProperty('content', content))
      ..add(DiagnosticsProperty('userId', userId))
      ..add(DiagnosticsProperty('createdAt', createdAt))
      ..add(DiagnosticsProperty('updatedAt', updatedAt))
      ..add(DiagnosticsProperty('user', user));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _Message &&
            (identical(other.id, id) ||
                const DeepCollectionEquality().equals(other.id, id)) &&
            (identical(other.content, content) ||
                const DeepCollectionEquality()
                    .equals(other.content, content)) &&
            (identical(other.userId, userId) ||
                const DeepCollectionEquality().equals(other.userId, userId)) &&
            (identical(other.createdAt, createdAt) ||
                const DeepCollectionEquality()
                    .equals(other.createdAt, createdAt)) &&
            (identical(other.updatedAt, updatedAt) ||
                const DeepCollectionEquality()
                    .equals(other.updatedAt, updatedAt)) &&
            (identical(other.user, user) ||
                const DeepCollectionEquality().equals(other.user, user)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^
      const DeepCollectionEquality().hash(id) ^
      const DeepCollectionEquality().hash(content) ^
      const DeepCollectionEquality().hash(userId) ^
      const DeepCollectionEquality().hash(createdAt) ^
      const DeepCollectionEquality().hash(updatedAt) ^
      const DeepCollectionEquality().hash(user);

  @JsonKey(ignore: true)
  @override
  _$MessageCopyWith<_Message> get copyWith =>
      __$MessageCopyWithImpl<_Message>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$_$_MessageToJson(this);
  }
}

abstract class _Message implements Message {
  const factory _Message(
      {String id,
      String content,
      String userId,
      DateTime createdAt,
      DateTime updatedAt,
      @JsonKey(name: 'user') User user}) = _$_Message;

  factory _Message.fromJson(Map<String, dynamic> json) = _$_Message.fromJson;

  @override
  String get id;
  @override
  String get content;
  @override
  String get userId;
  @override
  DateTime get createdAt;
  @override
  DateTime get updatedAt;
  @override
  @JsonKey(name: 'user')
  User get user;
  @override
  @JsonKey(ignore: true)
  _$MessageCopyWith<_Message> get copyWith;
}
