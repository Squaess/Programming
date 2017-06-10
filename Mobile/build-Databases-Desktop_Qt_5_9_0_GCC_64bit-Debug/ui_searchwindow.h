/********************************************************************************
** Form generated from reading UI file 'searchwindow.ui'
**
** Created by: Qt User Interface Compiler version 5.9.0
**
** WARNING! All changes made in this file will be lost when recompiling UI file!
********************************************************************************/

#ifndef UI_SEARCHWINDOW_H
#define UI_SEARCHWINDOW_H

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

class Ui_SearchWindow
{
public:
    QWidget *centralwidget;
    QVBoxLayout *verticalLayout;
    QHBoxLayout *horizontalLayout;
    QLabel *snL;
    QLineEdit *snF;
    QPushButton *snsearchB;
    QHBoxLayout *horizontalLayout_2;
    QLabel *srL;
    QLineEdit *srF;
    QPushButton *sssearchB;
    QHBoxLayout *horizontalLayout_3;
    QLabel *cnL;
    QLineEdit *crF;
    QPushButton *cnsearchB;

    void setupUi(QMainWindow *SearchWindow)
    {
        if (SearchWindow->objectName().isEmpty())
            SearchWindow->setObjectName(QStringLiteral("SearchWindow"));
        SearchWindow->resize(800, 600);
        centralwidget = new QWidget(SearchWindow);
        centralwidget->setObjectName(QStringLiteral("centralwidget"));
        verticalLayout = new QVBoxLayout(centralwidget);
        verticalLayout->setObjectName(QStringLiteral("verticalLayout"));
        horizontalLayout = new QHBoxLayout();
        horizontalLayout->setObjectName(QStringLiteral("horizontalLayout"));
        snL = new QLabel(centralwidget);
        snL->setObjectName(QStringLiteral("snL"));

        horizontalLayout->addWidget(snL);

        snF = new QLineEdit(centralwidget);
        snF->setObjectName(QStringLiteral("snF"));

        horizontalLayout->addWidget(snF);

        snsearchB = new QPushButton(centralwidget);
        snsearchB->setObjectName(QStringLiteral("snsearchB"));

        horizontalLayout->addWidget(snsearchB);


        verticalLayout->addLayout(horizontalLayout);

        horizontalLayout_2 = new QHBoxLayout();
        horizontalLayout_2->setObjectName(QStringLiteral("horizontalLayout_2"));
        srL = new QLabel(centralwidget);
        srL->setObjectName(QStringLiteral("srL"));

        horizontalLayout_2->addWidget(srL);

        srF = new QLineEdit(centralwidget);
        srF->setObjectName(QStringLiteral("srF"));

        horizontalLayout_2->addWidget(srF);

        sssearchB = new QPushButton(centralwidget);
        sssearchB->setObjectName(QStringLiteral("sssearchB"));

        horizontalLayout_2->addWidget(sssearchB);


        verticalLayout->addLayout(horizontalLayout_2);

        horizontalLayout_3 = new QHBoxLayout();
        horizontalLayout_3->setObjectName(QStringLiteral("horizontalLayout_3"));
        cnL = new QLabel(centralwidget);
        cnL->setObjectName(QStringLiteral("cnL"));

        horizontalLayout_3->addWidget(cnL);

        crF = new QLineEdit(centralwidget);
        crF->setObjectName(QStringLiteral("crF"));

        horizontalLayout_3->addWidget(crF);

        cnsearchB = new QPushButton(centralwidget);
        cnsearchB->setObjectName(QStringLiteral("cnsearchB"));

        horizontalLayout_3->addWidget(cnsearchB);


        verticalLayout->addLayout(horizontalLayout_3);

        SearchWindow->setCentralWidget(centralwidget);

        retranslateUi(SearchWindow);

        QMetaObject::connectSlotsByName(SearchWindow);
    } // setupUi

    void retranslateUi(QMainWindow *SearchWindow)
    {
        SearchWindow->setWindowTitle(QApplication::translate("SearchWindow", "MainWindow", Q_NULLPTR));
        snL->setText(QApplication::translate("SearchWindow", "Student Name:", Q_NULLPTR));
        snsearchB->setText(QApplication::translate("SearchWindow", "SEARCH", Q_NULLPTR));
        srL->setText(QApplication::translate("SearchWindow", "Student Surname", Q_NULLPTR));
        sssearchB->setText(QApplication::translate("SearchWindow", "SEARCH", Q_NULLPTR));
        cnL->setText(QApplication::translate("SearchWindow", "Course Name", Q_NULLPTR));
        cnsearchB->setText(QApplication::translate("SearchWindow", "SEARCH", Q_NULLPTR));
    } // retranslateUi

};

namespace Ui {
    class SearchWindow: public Ui_SearchWindow {};
} // namespace Ui

QT_END_NAMESPACE

#endif // UI_SEARCHWINDOW_H
