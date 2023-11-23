/* DEFINITIONS FILE MISSING ADDER TAGS*/
#include <defs_for_defs.h>


typedef struct
{
  REC_NO                CorrectDefinitionsChainStart;   /* if TM on Tms e.g pallet of totes */


/* PLUGINSTART (correct_definitions.CORRECT_DEFINITIONS_MEM_REC.plugin.inc)   */
/*pi*/    /* content from file: op.plugin... */
/*pi*/
/*pi*/    REC_NO                XxxNextTmRecNo;
/*pi*/
/* PLUGINEND - end of plugin - edit keyline do not alter                    */

  BOOLEAN Something; //TestComment
  REC_NO Yyy_ZzzLocNextTmRecNo;

} MISSING_ADDERTAGS_DEFINITIONS_MEM_REC_TYPE;

typedef struct
{
  ID_TYPE                    TmId;                   /* INDEX */
  CHAIN_TYPE               SelectedDestination;    /* (see above) */
  CHAIN_TYPE               RoutingDestination;     /* (see above) */
  REASON_TYPE              MoveReason;
  int                           ReasonRecNo;            /* needs to be int not REC_NO */

/* PLUGINSTART (tm_definitions.TM_REC_TYPE.plugin.inc)   */
/*pi*/    /* content from file: sorter_chute_loc.plugin... */
/*pi*/
/*pi*/    REC_NO   SorterChuteLocRecNo;
/*pi*/
/* PLUGINEND - end of plugin - edit keyline do not alter                    */


  REASON_TYPE ForceDespatchReqd;

  /* Used in RTV / STO orders */
  TIME_TYPE                 TimePalletLabelFirstPrinted;


  BOOLEAN Alberto; // test
  BOOLEAN Flag1; // hi
  BOOLEAN Flag1; // asdaksjhndkasbd

  /* DEFNONDBFLD int MyBool; */ // TESTCOMMENT
  /* DEFNONDBFLD int MyBool; */ // bleh

  /* DEFNONDBFLD int                       RetryButton; */

} MISSING_ADDERTAGS_DEFINITIONS_REC_TYPE;