/********************************************************************************
** Form generated from reading UI file 'deletewindow.ui'
**
** Created by: Qt User Interface Compiler version 5.9.0
**
** WARNING! All changes made in this file will be lost when recompiling UI file!
********************************************************************************/

#ifndef UI_DELETEWINDOW_H
#define UI_DELETEWINDOW_H

#include <QtCore/QVariant>
#include <QtWidgets/QAction>
#include <QtWidgets/QApplication>
#include <QtWidgets/QButtonGroup>
#include <QtWidgets/QHBoxLayout>
#include <QtWidgets/QHeaderView>
#include <QtWidgets/QLabel>
#include <QtWidgets/QLineEdit>
#include <QtWidgets/QMainWindow>
#include <QtWidgets/QPushButton>
#include <QtWidgets/QVBoxLayout>
#include <QtWidgets/QWidget>

QT_BEGIN_NAMESPACE

class Ui_DeleteWindow
{
public:
    QWidget *centralwidget;
    QHBoxLayout *horizontalLayout_3;
    QVBoxLayout *verticalLayout_3;
    QVBoxLayout *verticalLayout;
    QHBoxLayout *horizontalLayout;
    QLabel *sid;
    QLineEdit *sidF;
    QPushButton *pushButton;
    QVBoxLayout *verticalLayout_2;
    QHBoxLayout *horizontalLayout_2;
    QLabel *cid;
    QLineEdit *cidF;
    QPushButton *deleteB;

    void setupUi(QMainWindow *DeleteWindow)
    {
        if (DeleteWindow->objectName().isEmpty())
            DeleteWindow->setObjectName(QStringLiteral("DeleteWindow"));
        DeleteWindow->resize(800, 600);
        centralwidget = new QWidget(DeleteWindow);
        centralwidget->setObjectName(QStringLiteral("centralwidget"));
        horizontalLayout_3 = new QHBoxLayout(centralwidget);
        horizontalLayout_3->setObjectName(QStringLiteral("horizontalLayout_3"));
        verticalLayout_3 = new QVBoxLayout();
        verticalLayout_3->setObjectName(QStringLiteral("verticalLayout_3"));
        verticalLayout = new QVBoxLayout();
        verticalLayout->setObjectName(QStringLiteral("verticalLayout"));
        horizontalLayout = new QHBoxLayout();
        horizontalLayout->setObjectName(QStringLiteral("horizontalLayout"));
        sid = new QLabel(centralwidget);
        sid->setObjectName(QStringLiteral("sid"));

        horizontalLayout->addWidget(sid);

        sidF = new QLineEdit(centralwidget);
        sidF->setObjectName(QStringLiteral("sidF"));

        horizontalLayout->addWidget(sidF);


        verticalLayout->addLayout(horizontalLayout);

        pushButton = new QPushButton(centralwidget);
        pushButton->setObjectName(QStringLiteral("pushButton"));

        verticalLayout->addWidget(pushButton);


        verticalLayout_3->addLayout(verticalLayout);

        verticalLayout_2 = new QVBoxLayout();
        verticalLayout_2->setObjectName(QStringLiteral("verticalLayout_2"));
        horizontalLayout_2 = new QHBoxLayout();
        horizontalLayout_2->setObjectName(QStringLiteral("horizontalLayout_2"));
        cid = new QLabel(centralwidget);
        cid->setObjectName(QStringLiteral("cid"));

        horizontalLayout_2->addWidget(cid);

        cidF = new QLineEdit(centralwidget);
        cidF->setObjectName(QStringLiteral("cidF"));

        horizontalLayout_2->addWidget(cidF);


        verticalLayout_2->addLayout(horizontalLayout_2);

        deleteB = new QPushButton(centralwidget);
        deleteB->setObjectName(QStringLiteral("deleteB"));

        verticalLayout_2->addWidget(deleteB);


        verticalLayout_3->addLayout(verticalLayout_2);


        horizontalLayout_3->addLayout(verticalLayout_3);

        DeleteWindow->setCentralWidget(centralwidget);

        retranslateUi(DeleteWindow);

        QMetaObject::connectSlotsByName(DeleteWindow);
    } // setupUi

    void retranslateUi(QMainWindow *DeleteWindow)
    {
        DeleteWindow->setWindowTitle(QApplication::translate("DeleteWindow", "MainWindow", Q_NULLPTR));
        sid->setText(QApplication::translate("DeleteWindow", "st:id", Q_NULLPTR));
        pushButton->setText(QApplication::translate("DeleteWindow", "DELETE STUDENT", Q_NULLPTR));
        cid->setText(QApplication::translate("DeleteWindow", "c:id", Q_NULLPTR));
        deleteB->setText(QApplication::translate("DeleteWindow", "DELETE COURSE", Q_NULLPTR));
    } // retranslateUi

};

namespace Ui {
    class DeleteWindow: public Ui_DeleteWindow {};
} // namespace Ui

QT_END_NAMESPACE

#endif // UI_DELETEWINDOW_H
