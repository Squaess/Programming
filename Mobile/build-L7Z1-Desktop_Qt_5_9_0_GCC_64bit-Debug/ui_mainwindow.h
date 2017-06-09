/********************************************************************************
** Form generated from reading UI file 'mainwindow.ui'
**
** Created by: Qt User Interface Compiler version 5.9.0
**
** WARNING! All changes made in this file will be lost when recompiling UI file!
********************************************************************************/

#ifndef UI_MAINWINDOW_H
#define UI_MAINWINDOW_H

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

class Ui_MainWindow
{
public:
    QWidget *centralWidget;
    QHBoxLayout *horizontalLayout_4;
    QVBoxLayout *verticalLayout;
    QLabel *titleL;
    QHBoxLayout *horizontalLayout;
    QLabel *wordL;
    QLineEdit *lineEdit;
    QHBoxLayout *horizontalLayout_2;
    QLabel *correctL;
    QLabel *failedL;
    QHBoxLayout *horizontalLayout_3;
    QPushButton *nextButton;
    QPushButton *checkButton;
    QPushButton *showButton;

    void setupUi(QMainWindow *MainWindow)
    {
        if (MainWindow->objectName().isEmpty())
            MainWindow->setObjectName(QStringLiteral("MainWindow"));
        MainWindow->resize(430, 300);
        centralWidget = new QWidget(MainWindow);
        centralWidget->setObjectName(QStringLiteral("centralWidget"));
        horizontalLayout_4 = new QHBoxLayout(centralWidget);
        horizontalLayout_4->setSpacing(6);
        horizontalLayout_4->setContentsMargins(11, 11, 11, 11);
        horizontalLayout_4->setObjectName(QStringLiteral("horizontalLayout_4"));
        verticalLayout = new QVBoxLayout();
        verticalLayout->setSpacing(9);
        verticalLayout->setObjectName(QStringLiteral("verticalLayout"));
        verticalLayout->setContentsMargins(12, -1, 12, 3);
        titleL = new QLabel(centralWidget);
        titleL->setObjectName(QStringLiteral("titleL"));
        titleL->setAlignment(Qt::AlignCenter);

        verticalLayout->addWidget(titleL);

        horizontalLayout = new QHBoxLayout();
        horizontalLayout->setSpacing(12);
        horizontalLayout->setObjectName(QStringLiteral("horizontalLayout"));
        wordL = new QLabel(centralWidget);
        wordL->setObjectName(QStringLiteral("wordL"));
        wordL->setStyleSheet(QStringLiteral("font: 16pt;"));

        horizontalLayout->addWidget(wordL);

        lineEdit = new QLineEdit(centralWidget);
        lineEdit->setObjectName(QStringLiteral("lineEdit"));
        lineEdit->setMaximumSize(QSize(200, 16777215));
        lineEdit->setStyleSheet(QStringLiteral("font: 16pt;"));
        lineEdit->setAlignment(Qt::AlignBottom|Qt::AlignRight|Qt::AlignTrailing);

        horizontalLayout->addWidget(lineEdit);


        verticalLayout->addLayout(horizontalLayout);

        horizontalLayout_2 = new QHBoxLayout();
        horizontalLayout_2->setSpacing(12);
        horizontalLayout_2->setObjectName(QStringLiteral("horizontalLayout_2"));
        correctL = new QLabel(centralWidget);
        correctL->setObjectName(QStringLiteral("correctL"));
        correctL->setStyleSheet(QLatin1String("color: rgb(115, 210, 22);\n"
"font: 16pt ;"));

        horizontalLayout_2->addWidget(correctL);

        failedL = new QLabel(centralWidget);
        failedL->setObjectName(QStringLiteral("failedL"));
        failedL->setStyleSheet(QLatin1String("color: rgb(164, 0, 0);\n"
"font: 16pt;"));
        failedL->setAlignment(Qt::AlignRight|Qt::AlignTrailing|Qt::AlignVCenter);

        horizontalLayout_2->addWidget(failedL);


        verticalLayout->addLayout(horizontalLayout_2);

        horizontalLayout_3 = new QHBoxLayout();
        horizontalLayout_3->setSpacing(6);
        horizontalLayout_3->setObjectName(QStringLiteral("horizontalLayout_3"));
        nextButton = new QPushButton(centralWidget);
        nextButton->setObjectName(QStringLiteral("nextButton"));

        horizontalLayout_3->addWidget(nextButton);

        checkButton = new QPushButton(centralWidget);
        checkButton->setObjectName(QStringLiteral("checkButton"));

        horizontalLayout_3->addWidget(checkButton);

        showButton = new QPushButton(centralWidget);
        showButton->setObjectName(QStringLiteral("showButton"));

        horizontalLayout_3->addWidget(showButton);


        verticalLayout->addLayout(horizontalLayout_3);


        horizontalLayout_4->addLayout(verticalLayout);

        MainWindow->setCentralWidget(centralWidget);
#ifndef QT_NO_SHORTCUT
        wordL->setBuddy(lineEdit);
#endif // QT_NO_SHORTCUT

        retranslateUi(MainWindow);

        QMetaObject::connectSlotsByName(MainWindow);
    } // setupUi

    void retranslateUi(QMainWindow *MainWindow)
    {
        MainWindow->setWindowTitle(QApplication::translate("MainWindow", "MainWindow", Q_NULLPTR));
        titleL->setText(QApplication::translate("MainWindow", "Guess from englidh to polish", Q_NULLPTR));
        wordL->setText(QApplication::translate("MainWindow", "home", Q_NULLPTR));
        correctL->setText(QApplication::translate("MainWindow", "<html><head/><body><p><span style=\" color:#73d216;\">Correct: -</span></p></body></html>", Q_NULLPTR));
        failedL->setText(QApplication::translate("MainWindow", "<html><head/><body><p><span style=\" color:#cc0000;\">Failed: -</span></p></body></html>", Q_NULLPTR));
        nextButton->setText(QApplication::translate("MainWindow", "Next", Q_NULLPTR));
        checkButton->setText(QApplication::translate("MainWindow", "Check", Q_NULLPTR));
        showButton->setText(QApplication::translate("MainWindow", "Show", Q_NULLPTR));
    } // retranslateUi

};

namespace Ui {
    class MainWindow: public Ui_MainWindow {};
} // namespace Ui

QT_END_NAMESPACE

#endif // UI_MAINWINDOW_H
