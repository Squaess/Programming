/****************************************************************************
** Meta object code from reading C++ file 'insertwindow.h'
**
** Created by: The Qt Meta Object Compiler version 67 (Qt 5.9.0)
**
** WARNING! All changes made in this file will be lost!
*****************************************************************************/

#include "../Databases/insertwindow.h"
#include <QtCore/qbytearray.h>
#include <QtCore/qmetatype.h>
#if !defined(Q_MOC_OUTPUT_REVISION)
#error "The header file 'insertwindow.h' doesn't include <QObject>."
#elif Q_MOC_OUTPUT_REVISION != 67
#error "This file was generated using the moc from 5.9.0. It"
#error "cannot be used with the include files from this version of Qt."
#error "(The moc has changed too much.)"
#endif

QT_BEGIN_MOC_NAMESPACE
QT_WARNING_PUSH
QT_WARNING_DISABLE_DEPRECATED
struct qt_meta_stringdata_InsertWindow_t {
    QByteArrayData data[9];
    char stringdata0[105];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_InsertWindow_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_InsertWindow_t qt_meta_stringdata_InsertWindow = {
    {
QT_MOC_LITERAL(0, 0, 12), // "InsertWindow"
QT_MOC_LITERAL(1, 13, 10), // "plsAddStud"
QT_MOC_LITERAL(2, 24, 0), // ""
QT_MOC_LITERAL(3, 25, 4), // "data"
QT_MOC_LITERAL(4, 30, 11), // "plsAddGrade"
QT_MOC_LITERAL(5, 42, 12), // "plsAddCourse"
QT_MOC_LITERAL(6, 55, 15), // "on_addB_clicked"
QT_MOC_LITERAL(7, 71, 16), // "on_addGB_clicked"
QT_MOC_LITERAL(8, 88, 16) // "on_addCB_clicked"

    },
    "InsertWindow\0plsAddStud\0\0data\0plsAddGrade\0"
    "plsAddCourse\0on_addB_clicked\0"
    "on_addGB_clicked\0on_addCB_clicked"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_InsertWindow[] = {

 // content:
       7,       // revision
       0,       // classname
       0,    0, // classinfo
       6,   14, // methods
       0,    0, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       3,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    1,   44,    2, 0x06 /* Public */,
       4,    1,   47,    2, 0x06 /* Public */,
       5,    1,   50,    2, 0x06 /* Public */,

 // slots: name, argc, parameters, tag, flags
       6,    0,   53,    2, 0x08 /* Private */,
       7,    0,   54,    2, 0x08 /* Private */,
       8,    0,   55,    2, 0x08 /* Private */,

 // signals: parameters
    QMetaType::Void, QMetaType::QString,    3,
    QMetaType::Void, QMetaType::QString,    3,
    QMetaType::Void, QMetaType::QString,    3,

 // slots: parameters
    QMetaType::Void,
    QMetaType::Void,
    QMetaType::Void,

       0        // eod
};

void InsertWindow::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        InsertWindow *_t = static_cast<InsertWindow *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->plsAddStud((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 1: _t->plsAddGrade((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 2: _t->plsAddCourse((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 3: _t->on_addB_clicked(); break;
        case 4: _t->on_addGB_clicked(); break;
        case 5: _t->on_addCB_clicked(); break;
        default: ;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        void **func = reinterpret_cast<void **>(_a[1]);
        {
            typedef void (InsertWindow::*_t)(QString );
            if (*reinterpret_cast<_t *>(func) == static_cast<_t>(&InsertWindow::plsAddStud)) {
                *result = 0;
                return;
            }
        }
        {
            typedef void (InsertWindow::*_t)(QString );
            if (*reinterpret_cast<_t *>(func) == static_cast<_t>(&InsertWindow::plsAddGrade)) {
                *result = 1;
                return;
            }
        }
        {
            typedef void (InsertWindow::*_t)(QString );
            if (*reinterpret_cast<_t *>(func) == static_cast<_t>(&InsertWindow::plsAddCourse)) {
                *result = 2;
                return;
            }
        }
    }
}

const QMetaObject InsertWindow::staticMetaObject = {
    { &QMainWindow::staticMetaObject, qt_meta_stringdata_InsertWindow.data,
      qt_meta_data_InsertWindow,  qt_static_metacall, nullptr, nullptr}
};


const QMetaObject *InsertWindow::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *InsertWindow::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_InsertWindow.stringdata0))
        return static_cast<void*>(const_cast< InsertWindow*>(this));
    return QMainWindow::qt_metacast(_clname);
}

int InsertWindow::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QMainWindow::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 6)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 6;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 6)
            *reinterpret_cast<int*>(_a[0]) = -1;
        _id -= 6;
    }
    return _id;
}

// SIGNAL 0
void InsertWindow::plsAddStud(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}

// SIGNAL 1
void InsertWindow::plsAddGrade(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}

// SIGNAL 2
void InsertWindow::plsAddCourse(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}
QT_WARNING_POP
QT_END_MOC_NAMESPACE
