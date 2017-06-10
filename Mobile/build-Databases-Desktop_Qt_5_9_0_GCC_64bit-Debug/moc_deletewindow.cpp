/****************************************************************************
** Meta object code from reading C++ file 'deletewindow.h'
**
** Created by: The Qt Meta Object Compiler version 67 (Qt 5.9.0)
**
** WARNING! All changes made in this file will be lost!
*****************************************************************************/

#include "../Databases/deletewindow.h"
#include <QtCore/qbytearray.h>
#include <QtCore/qmetatype.h>
#if !defined(Q_MOC_OUTPUT_REVISION)
#error "The header file 'deletewindow.h' doesn't include <QObject>."
#elif Q_MOC_OUTPUT_REVISION != 67
#error "This file was generated using the moc from 5.9.0. It"
#error "cannot be used with the include files from this version of Qt."
#error "(The moc has changed too much.)"
#endif

QT_BEGIN_MOC_NAMESPACE
QT_WARNING_PUSH
QT_WARNING_DISABLE_DEPRECATED
struct qt_meta_stringdata_DeleteWindow_t {
    QByteArrayData data[7];
    char stringdata0[93];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_DeleteWindow_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_DeleteWindow_t qt_meta_stringdata_DeleteWindow = {
    {
QT_MOC_LITERAL(0, 0, 12), // "DeleteWindow"
QT_MOC_LITERAL(1, 13, 16), // "plsDeleteStudent"
QT_MOC_LITERAL(2, 30, 0), // ""
QT_MOC_LITERAL(3, 31, 4), // "data"
QT_MOC_LITERAL(4, 36, 15), // "plsDeleteCourse"
QT_MOC_LITERAL(5, 52, 21), // "on_pushButton_clicked"
QT_MOC_LITERAL(6, 74, 18) // "on_deleteB_clicked"

    },
    "DeleteWindow\0plsDeleteStudent\0\0data\0"
    "plsDeleteCourse\0on_pushButton_clicked\0"
    "on_deleteB_clicked"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_DeleteWindow[] = {

 // content:
       7,       // revision
       0,       // classname
       0,    0, // classinfo
       4,   14, // methods
       0,    0, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       2,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    1,   34,    2, 0x06 /* Public */,
       4,    1,   37,    2, 0x06 /* Public */,

 // slots: name, argc, parameters, tag, flags
       5,    0,   40,    2, 0x08 /* Private */,
       6,    0,   41,    2, 0x08 /* Private */,

 // signals: parameters
    QMetaType::Void, QMetaType::QString,    3,
    QMetaType::Void, QMetaType::QString,    3,

 // slots: parameters
    QMetaType::Void,
    QMetaType::Void,

       0        // eod
};

void DeleteWindow::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        DeleteWindow *_t = static_cast<DeleteWindow *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->plsDeleteStudent((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 1: _t->plsDeleteCourse((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 2: _t->on_pushButton_clicked(); break;
        case 3: _t->on_deleteB_clicked(); break;
        default: ;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        void **func = reinterpret_cast<void **>(_a[1]);
        {
            typedef void (DeleteWindow::*_t)(QString );
            if (*reinterpret_cast<_t *>(func) == static_cast<_t>(&DeleteWindow::plsDeleteStudent)) {
                *result = 0;
                return;
            }
        }
        {
            typedef void (DeleteWindow::*_t)(QString );
            if (*reinterpret_cast<_t *>(func) == static_cast<_t>(&DeleteWindow::plsDeleteCourse)) {
                *result = 1;
                return;
            }
        }
    }
}

const QMetaObject DeleteWindow::staticMetaObject = {
    { &QMainWindow::staticMetaObject, qt_meta_stringdata_DeleteWindow.data,
      qt_meta_data_DeleteWindow,  qt_static_metacall, nullptr, nullptr}
};


const QMetaObject *DeleteWindow::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *DeleteWindow::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_DeleteWindow.stringdata0))
        return static_cast<void*>(const_cast< DeleteWindow*>(this));
    return QMainWindow::qt_metacast(_clname);
}

int DeleteWindow::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QMainWindow::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 4)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 4;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 4)
            *reinterpret_cast<int*>(_a[0]) = -1;
        _id -= 4;
    }
    return _id;
}

// SIGNAL 0
void DeleteWindow::plsDeleteStudent(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}

// SIGNAL 1
void DeleteWindow::plsDeleteCourse(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}
QT_WARNING_POP
QT_END_MOC_NAMESPACE
