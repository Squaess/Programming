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
    QHBoxLayout *horizontalLayout_2;
    QVBoxLayout *verticalLayout;
    QGridLayout *gridLayout;
    QLabel *surnameL;
    QLineEdit *nameF;
    QLineEdit *surnameF;
    QLabel *nameL;
    QHBoxLayout *horizontalLayout;
    QPushButton *addB;
    QPushButton *cancelB;
    QStatusBar *statusbar;

    void setupUi(QMainWindow *InsertWindow)
    {
        if (InsertWindow->objectName().isEmpty())
            InsertWindow->setObjectName(QStringLiteral("InsertWindow"));
        InsertWindow->resize(800, 600);
        centralwidget = new QWidget(InsertWindow);
        centralwidget->setObjectName(QStringLiteral("centralwidget"));
        horizontalLayout_2 = new QHBoxLayout(centralwidget);
        horizontalLayout_2->setObjectName(QStringLiteral("horizontalLayout_2"));
        verticalLayout = new QVBoxLayout();
        verticalLayout->setObjectName(QStringLiteral("verticalLayout"));
        gridLayout = new QGridLayout();
        gridLayout->setObjectName(QStringLiteral("gridLayout"));
        gridLayout->setHorizontalSpacing(12);
        surnameL = new QLabel(centralwidget);
        surnameL->setObjectName(QStringLiteral("surnameL"));

        gridLayout->addWidget(surnameL, 1, 0, 1, 1);

        nameF = new QLineEdit(centralwidget);
        nameF->setObjectName(QStringLiteral("nameF"));

        gridLayout->addWidget(nameF, 0, 1, 1, 1);

        surnameF = new QLineEdit(centralwidget);
        surnameF->setObjectName(QStringLiteral("surnameF"));

        gridLayout->addWidget(surnameF, 1, 1, 1, 1);

        nameL = new QLabel(centralwidget);
        nameL->setObjectName(QStringLiteral("nameL"));

        gridLayout->addWidget(nameL, 0, 0, 1, 1);


        verticalLayout->addLayout(gridLayout);

        horizontalLayout = new QHBoxLayout();
        horizontalLayout->setObjectName(QStringLiteral("horizontalLayout"));
        addB = new QPushButton(centralwidget);
        addB->setObjectName(QStringLiteral("addB"));

        horizontalLayout->addWidget(addB);

        cancelB = new QPushButton(centralwidget);
        cancelB->setObjectName(QStringLiteral("cancelB"));

        horizontalLayout->addWidget(cancelB);


        verticalLayout->addLayout(horizontalLayout);


        horizontalLayout_2->addLayout(verticalLayout);

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
        surnameL->setText(QApplication::translate("InsertWindow", "SURNAME:", Q_NULLPTR));
        nameL->setText(QApplication::translate("InsertWindow", "NAME:", Q_NULLPTR));
        addB->setText(QApplication::translate("InsertWindow", "ADD", Q_NULLPTR));
        cancelB->setText(QApplication::translate("InsertWindow", "CANCEL", Q_NULLPTR));
    } // retranslateUi

};

namespace Ui {
    class InsertWindow: public Ui_InsertWindow {};
} // namespace Ui

QT_END_NAMESPACE

#endif // UI_INSERTWINDOW_H
