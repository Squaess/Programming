/****************************************************************************
** Meta object code from reading C++ file 'mainwindow.h'
**
** Created by: The Qt Meta Object Compiler version 67 (Qt 5.9.0)
**
** WARNING! All changes made in this file will be lost!
*****************************************************************************/

#include "../Databases/mainwindow.h"
#include <QtCore/qbytearray.h>
#include <QtCore/qmetatype.h>
#if !defined(Q_MOC_OUTPUT_REVISION)
#error "The header file 'mainwindow.h' doesn't include <QObject>."
#elif Q_MOC_OUTPUT_REVISION != 67
#error "This file was generated using the moc from 5.9.0. It"
#error "cannot be used with the include files from this version of Qt."
#error "(The moc has changed too much.)"
#endif

QT_BEGIN_MOC_NAMESPACE
QT_WARNING_PUSH
QT_WARNING_DISABLE_DEPRECATED
struct qt_meta_stringdata_MainWindow_t {
    QByteArrayData data[17];
    char stringdata0[220];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_MainWindow_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_MainWindow_t qt_meta_stringdata_MainWindow = {
    {
QT_MOC_LITERAL(0, 0, 10), // "MainWindow"
QT_MOC_LITERAL(1, 11, 16), // "on_showB_clicked"
QT_MOC_LITERAL(2, 28, 0), // ""
QT_MOC_LITERAL(3, 29, 18), // "on_insertB_clicked"
QT_MOC_LITERAL(4, 48, 14), // "on_ssB_clicked"
QT_MOC_LITERAL(5, 63, 14), // "on_scB_clicked"
QT_MOC_LITERAL(6, 78, 18), // "on_deleteB_clicked"
QT_MOC_LITERAL(7, 97, 18), // "on_searchB_clicked"
QT_MOC_LITERAL(8, 116, 10), // "addStudent"
QT_MOC_LITERAL(9, 127, 4), // "data"
QT_MOC_LITERAL(10, 132, 8), // "addGrade"
QT_MOC_LITERAL(11, 141, 9), // "addCourse"
QT_MOC_LITERAL(12, 151, 13), // "deleteStudent"
QT_MOC_LITERAL(13, 165, 12), // "deleteCourse"
QT_MOC_LITERAL(14, 178, 14), // "searchStudName"
QT_MOC_LITERAL(15, 193, 13), // "searchStudSur"
QT_MOC_LITERAL(16, 207, 12) // "searchCourse"

    },
    "MainWindow\0on_showB_clicked\0\0"
    "on_insertB_clicked\0on_ssB_clicked\0"
    "on_scB_clicked\0on_deleteB_clicked\0"
    "on_searchB_clicked\0addStudent\0data\0"
    "addGrade\0addCourse\0deleteStudent\0"
    "deleteCourse\0searchStudName\0searchStudSur\0"
    "searchCourse"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_MainWindow[] = {

 // content:
       7,       // revision
       0,       // classname
       0,    0, // classinfo
      14,   14, // methods
       0,    0, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       0,       // signalCount

 // slots: name, argc, parameters, tag, flags
       1,    0,   84,    2, 0x08 /* Private */,
       3,    0,   85,    2, 0x08 /* Private */,
       4,    0,   86,    2, 0x08 /* Private */,
       5,    0,   87,    2, 0x08 /* Private */,
       6,    0,   88,    2, 0x08 /* Private */,
       7,    0,   89,    2, 0x08 /* Private */,
       8,    1,   90,    2, 0x0a /* Public */,
      10,    1,   93,    2, 0x0a /* Public */,
      11,    1,   96,    2, 0x0a /* Public */,
      12,    1,   99,    2, 0x0a /* Public */,
      13,    1,  102,    2, 0x0a /* Public */,
      14,    1,  105,    2, 0x0a /* Public */,
      15,    1,  108,    2, 0x0a /* Public */,
      16,    1,  111,    2, 0x0a /* Public */,

 // slots: parameters
    QMetaType::Void,
    QMetaType::Void,
    QMetaType::Void,
    QMetaType::Void,
    QMetaType::Void,
    QMetaType::Void,
    QMetaType::Void, QMetaType::QString,    9,
    QMetaType::Void, QMetaType::QString,    9,
    QMetaType::Void, QMetaType::QString,    9,
    QMetaType::Void, QMetaType::QString,    9,
    QMetaType::Void, QMetaType::QString,    9,
    QMetaType::Void, QMetaType::QString,    9,
    QMetaType::Void, QMetaType::QString,    9,
    QMetaType::Void, QMetaType::QString,    9,

       0        // eod
};

void MainWindow::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        MainWindow *_t = static_cast<MainWindow *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->on_showB_clicked(); break;
        case 1: _t->on_insertB_clicked(); break;
        case 2: _t->on_ssB_clicked(); break;
        case 3: _t->on_scB_clicked(); break;
        case 4: _t->on_deleteB_clicked(); break;
        case 5: _t->on_searchB_clicked(); break;
        case 6: _t->addStudent((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 7: _t->addGrade((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 8: _t->addCourse((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 9: _t->deleteStudent((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 10: _t->deleteCourse((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 11: _t->searchStudName((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 12: _t->searchStudSur((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 13: _t->searchCourse((*reinterpret_cast< QString(*)>(_a[1]))); break;
        default: ;
        }
    }
}

const QMetaObject MainWindow::staticMetaObject = {
    { &QMainWindow::staticMetaObject, qt_meta_stringdata_MainWindow.data,
      qt_meta_data_MainWindow,  qt_static_metacall, nullptr, nullptr}
};


const QMetaObject *MainWindow::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *MainWindow::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_MainWindow.stringdata0))
        return static_cast<void*>(const_cast< MainWindow*>(this));
    return QMainWindow::qt_metacast(_clname);
}

int MainWindow::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QMainWindow::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 14)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 14;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 14)
            *reinterpret_cast<int*>(_a[0]) = -1;
        _id -= 14;
    }
    return _id;
}
QT_WARNING_POP
QT_END_MOC_NAMESPACE
