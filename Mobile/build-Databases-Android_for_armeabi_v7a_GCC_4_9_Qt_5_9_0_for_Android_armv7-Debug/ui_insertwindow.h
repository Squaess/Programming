/********************************************************************************
** Form generated from reading UI file 'insertwindow.ui'
**
** Created by: Qt User Interface Compiler version 5.9.0
**
** WARNING! All changes made in this file will be lost when recompiling UI file!
********************************************************************************/

#ifndef UI_INSERTWINDOW_H
#define UI_INSERTWINDOW_H

#include <QtCore/QVariant>
#include <QtWidgets/QAction>
#include <QtWidgets/QApplication>
#include <QtWidgets/QButtonGroup>
#include <QtWidgets/QGridLayout>
#include <QtWidgets/QHBoxLayout>
#include <QtWidgets/QHeaderView>
#include <QtWidgets/QLabel>
#include <QtWidgets/QLineEdit>
#include <QtWidgets/QMainWindow>
#include <QtWidgets/QPushButton>
#include <QtWidgets/QStatusBar>
#include <QtWidgets/QVBoxLayout>
#include <QtWidgets/QWidget>

QT_BEGIN_NAMESPACE

class Ui_InsertWindow
{
public:
    QWidget *centralwidget;
    QHBoxLayout *horizontalLayout;
    QVBoxLayout *verticalLayout_4;
    QVBoxLayout *verticalLayout;
    QGridLayout *gridLayout;
    QLabel *nameL;
    QLineEdit *nameF;
    QLabel *surnameL;
    QLineEdit *surnameF;
    QPushButton *addB;
    QVBoxLayout *verticalLayout_2;
    QGridLayout *gridLayout_2;
    QLabel *sid;
    QLineEdit *sidF;
    QLabel *cid;
    QLineEdit *cidF;
    QLabel *gid;
    QLineEdit *gidF;
    QPushButton *addGB;
    QVBoxLayout *verticalLayout_3;
    QHBoxLayout *horizontalLayout_2;
    QLabel *cnameL;
    QLineEdit *cnameF;
    QPushButton *addCB;
    QStatusBar *statusbar;

    void setupUi(QMainWindow *InsertWindow)
    {
        if (InsertWindow->objectName().isEmpty())
            InsertWindow->setObjectName(QStringLiteral("InsertWindow"));
        InsertWindow->resize(800, 600);
        centralwidget = new QWidget(InsertWindow);
        centralwidget->setObjectName(QStringLiteral("centralwidget"));
        horizontalLayout = new QHBoxLayout(centralwidget);
        horizontalLayout->setObjectName(QStringLiteral("horizontalLayout"));
        verticalLayout_4 = new QVBoxLayout();
        verticalLayout_4->setObjectName(QStringLiteral("verticalLayout_4"));
        verticalLayout = new QVBoxLayout();
        verticalLayout->setObjectName(QStringLiteral("verticalLayout"));
        gridLayout = new QGridLayout();
        gridLayout->setObjectName(QStringLiteral("gridLayout"));
        nameL = new QLabel(centralwidget);
        nameL->setObjectName(QStringLiteral("nameL"));

        gridLayout->addWidget(nameL, 0, 0, 1, 1);

        nameF = new QLineEdit(centralwidget);
        nameF->setObjectName(QStringLiteral("nameF"));

        gridLayout->addWidget(nameF, 0, 1, 1, 1);

        surnameL = new QLabel(centralwidget);
        surnameL->setObjectName(QStringLiteral("surnameL"));

        gridLayout->addWidget(surnameL, 1, 0, 1, 1);

        surnameF = new QLineEdit(centralwidget);
        surnameF->setObjectName(QStringLiteral("surnameF"));

        gridLayout->addWidget(surnameF, 1, 1, 1, 1);


        verticalLayout->addLayout(gridLayout);

        addB = new QPushButton(centralwidget);
        addB->setObjectName(QStringLiteral("addB"));

        verticalLayout->addWidget(addB);


        verticalLayout_4->addLayout(verticalLayout);

        verticalLayout_2 = new QVBoxLayout();
        verticalLayout_2->setObjectName(QStringLiteral("verticalLayout_2"));
        gridLayout_2 = new QGridLayout();
        gridLayout_2->setObjectName(QStringLiteral("gridLayout_2"));
        sid = new QLabel(centralwidget);
        sid->setObjectName(QStringLiteral("sid"));

        gridLayout_2->addWidget(sid, 0, 0, 1, 1);

        sidF = new QLineEdit(centralwidget);
        sidF->setObjectName(QStringLiteral("sidF"));

        gridLayout_2->addWidget(sidF, 0, 1, 1, 1);

        cid = new QLabel(centralwidget);
        cid->setObjectName(QStringLiteral("cid"));

        gridLayout_2->addWidget(cid, 1, 0, 1, 1);

        cidF = new QLineEdit(centralwidget);
        cidF->setObjectName(QStringLiteral("cidF"));

        gridLayout_2->addWidget(cidF, 1, 1, 1, 1);

        gid = new QLabel(centralwidget);
        gid->setObjectName(QStringLiteral("gid"));

        gridLayout_2->addWidget(gid, 2, 0, 1, 1);

        gidF = new QLineEdit(centralwidget);
        gidF->setObjectName(QStringLiteral("gidF"));

        gridLayout_2->addWidget(gidF, 2, 1, 1, 1);


        verticalLayout_2->addLayout(gridLayout_2);

        addGB = new QPushButton(centralwidget);
        addGB->setObjectName(QStringLiteral("addGB"));

        verticalLayout_2->addWidget(addGB);


        verticalLayout_4->addLayout(verticalLayout_2);

        verticalLayout_3 = new QVBoxLayout();
        verticalLayout_3->setObjectName(QStringLiteral("verticalLayout_3"));
        horizontalLayout_2 = new QHBoxLayout();
        horizontalLayout_2->setObjectName(QStringLiteral("horizontalLayout_2"));
        cnameL = new QLabel(centralwidget);
        cnameL->setObjectName(QStringLiteral("cnameL"));

        horizontalLayout_2->addWidget(cnameL);

        cnameF = new QLineEdit(centralwidget);
        cnameF->setObjectName(QStringLiteral("cnameF"));

        horizontalLayout_2->addWidget(cnameF);


        verticalLayout_3->addLayout(horizontalLayout_2);

        addCB = new QPushButton(centralwidget);
        addCB->setObjectName(QStringLiteral("addCB"));

        verticalLayout_3->addWidget(addCB);


        verticalLayout_4->addLayout(verticalLayout_3);


        horizontalLayout->addLayout(verticalLayout_4);

        InsertWindow->setCentralWidget(centralwidget);
        statusbar = new QStatusBar(InsertWindow);
        statusbar->setObjectName(QStringLiteral("statusbar"));
        InsertWindow->setStatusBar(statusbar);

        retranslateUi(InsertWindow);

        QMetaObject::connectSlotsByName(InsertWindow);
    } // setupUi

    void retranslateUi(QMainWindow *InsertWindow)
    {
        InsertWindow->setWindowTitle(QApplication::translate("InsertWindow", "MainWindow", Q_NULLPTR));
        nameL->setText(QApplication::translate("InsertWindow", "NAME:", Q_NULLPTR));
        surnameL->setText(QApplication::translate("InsertWindow", "SURNAME:", Q_NULLPTR));
        addB->setText(QApplication::translate("InsertWindow", "ADD STUDENT", Q_NULLPTR));
        sid->setText(QApplication::translate("InsertWindow", "ST:ID", Q_NULLPTR));
        cid->setText(QApplication::translate("InsertWindow", "C:ID", Q_NULLPTR));
        gid->setText(QApplication::translate("InsertWindow", "GRADE", Q_NULLPTR));
        addGB->setText(QApplication::translate("InsertWindow", "ADD GRADE", Q_NULLPTR));
        cnameL->setText(QApplication::translate("InsertWindow", "COURSE NAME:", Q_NULLPTR));
        addCB->setText(QApplication::translate("InsertWindow", "ADD COURSE", Q_NULLPTR));
    } // retranslateUi

};

namespace Ui {
    class InsertWindow: public Ui_InsertWindow {};
} // namespace Ui

QT_END_NAMESPACE

#endif // UI_INSERTWINDOW_H
